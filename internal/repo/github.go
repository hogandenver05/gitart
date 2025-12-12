package repo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type PushStatus struct {
	RepositoryName     string
	RepositoryPath     string
	Username           string
	RemoteURL          string
	Branch             string
	RepoAlreadyExists  bool
}

// ArtworkRegenerator regenerates artwork commits in the repository after a reset.
type ArtworkRegenerator func() error

// PushToGitHub deletes any existing remote repository, optionally resets the local repository,
// regenerates artwork if reset, and pushes to a new GitHub repository.
func (repository *NestedRepository) PushToGitHub(private bool, reset bool, regenerateArtwork ArtworkRegenerator) (*PushStatus, error) {
	repositoryName := "gitart-" + time.Now().Format("2006-01-02")
	repositoryPath, err := filepath.Abs(repository.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path: %w", err)
	}

	username := GetUsername()
	remoteURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repositoryName)

	status := &PushStatus{
		RepositoryName: repositoryName,
		RepositoryPath: repositoryPath,
		Username:       username,
		RemoteURL:      remoteURL,
	}

	repoAlreadyExists, err := RepositoryExists(repositoryName)
	if err != nil {
		return nil, fmt.Errorf("failed to check repository existence: %w", err)
	}

	if repoAlreadyExists {
		if err := DeleteRepository(repositoryName); err != nil {
			return nil, fmt.Errorf("failed to delete existing repository: %w", err)
		}
		status.RepoAlreadyExists = true
	}

	if reset {
		if err := repository.ResetRepository(); err != nil {
			return nil, fmt.Errorf("failed to reset repository: %w", err)
		}
		if regenerateArtwork != nil {
			if err := regenerateArtwork(); err != nil {
				return nil, fmt.Errorf("failed to regenerate artwork after reset: %w", err)
			}
		}
		if err := repository.IncludeREADMEIfPresent(); err != nil {
			return nil, fmt.Errorf("failed to include README after reset: %w", err)
		}
	}

	if err := CreateRepository(repositoryName, repositoryPath, private); err != nil {
		return nil, fmt.Errorf("failed to create GitHub repository: %w", err)
	}

	if err := AddRemoteOrigin(repositoryName, repositoryPath); err != nil {
		return nil, fmt.Errorf("failed to set remote origin: %w", err)
	}

	defaultBranch, err := getDefaultBranch(repositoryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to detect default branch: %w", err)
	}
	status.Branch = defaultBranch

	cmd := exec.Command("git", "push", "--force", "-u", "origin", defaultBranch)
	cmd.Dir = repositoryPath
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to push commits: %w", err)
	}

	return status, nil
}

// GetUsername retrieves the current GitHub username using the GitHub CLI.
func GetUsername() string {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// RepositoryExists checks if a GitHub repository exists for the current user.
func RepositoryExists(repositoryName string) (bool, error) {
	cmd := exec.Command("gh", "repo", "view", repositoryName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputText := string(output)
		if strings.Contains(outputText, "not found") || strings.Contains(outputText, "Could not resolve") {
			return false, nil
		}
		return false, fmt.Errorf("failed to check repository: %w, output: %s", err, outputText)
	}
	return true, nil
}

// DeleteRepository deletes a GitHub repository. Returns nil if the repository does not exist.
func DeleteRepository(repositoryName string) error {
	cmd := exec.Command("gh", "repo", "delete", repositoryName, "--yes")
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputText := string(output)
		if strings.Contains(outputText, "not found") || strings.Contains(outputText, "Could not resolve") {
			return nil
		}
		return fmt.Errorf("failed to delete repository: %w, output: %s", err, outputText)
	}
	return nil
}

// CreateRepository creates a new GitHub repository with the specified name and visibility.
func CreateRepository(repositoryName string, dir string, private bool) error {
	args := []string{"repo", "create", repositoryName}
	if private {
		args = append(args, "--private")
	} else {
		args = append(args, "--public")
	}

	cmd := exec.Command("gh", args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputText := string(output)
		if strings.Contains(outputText, "already exists") {
			return fmt.Errorf("repository already exists (should have been deleted): %s", repositoryName)
		}
		return fmt.Errorf("failed to create GitHub repo: %w, output: %s", err, outputText)
	}

	return nil
}

// ResetRepository removes the existing git repository and initializes a new one.
func (repository *NestedRepository) ResetRepository() error {
	gitPath := filepath.Join(repository.Path, ".git")
	if _, err := os.Stat(gitPath); err == nil {
		if err := os.RemoveAll(gitPath); err != nil {
			return fmt.Errorf("failed to remove existing .git folder: %w", err)
		}
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = repository.Path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	return nil
}

// AddRemoteOrigin adds or updates the origin remote for the repository.
func AddRemoteOrigin(repositoryName string, dir string) error {
	username := GetUsername()
	url := fmt.Sprintf("https://github.com/%s/%s.git", username, repositoryName)

	checkCmd := exec.Command("git", "remote", "get-url", "origin")
	checkCmd.Dir = dir
	if err := checkCmd.Run(); err == nil {
		cmd := exec.Command("git", "remote", "set-url", "origin", url)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to update remote: %w", err)
		}
	} else {
		cmd := exec.Command("git", "remote", "add", "origin", url)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to add remote: %w", err)
		}
	}

	return nil
}

// getDefaultBranch returns the name of the default branch in the repository.
func getDefaultBranch(dir string) (string, error) {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to list branches: %w", err)
	}

	branches := strings.Fields(strings.ReplaceAll(string(out), "*", ""))
	if len(branches) == 0 {
		return "", fmt.Errorf("no branches found")
	}

	return branches[0], nil
}
