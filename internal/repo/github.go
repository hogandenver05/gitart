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

func (repository *NestedRepository) PushToGitHub(private bool, reset bool) (*PushStatus, error) {
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

	repoAlreadyExists, err := CreateRepository(repositoryName, repositoryPath, private)
	if err != nil {
		return nil, fmt.Errorf("failed to create GitHub repository: %w", err)
	}
	status.RepoAlreadyExists = repoAlreadyExists

	if repoAlreadyExists && reset {
		if err := repository.ResetRepository(); err != nil {
			return nil, fmt.Errorf("failed to reset repository: %w", err)
		}
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

func GetUsername() string {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func CreateRepository(repositoryName string, dir string, private bool) (bool, error) {
	args := []string{"repo", "create", repositoryName}
	if private {
		args = append(args, "--private")
	} else {
		args = append(args, "--public")
	}

	cmd := exec.Command("gh", args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	outputText := string(output)

	if err != nil {
		if strings.Contains(outputText, "already exists") {
			return true, nil
		}
		return false, fmt.Errorf("failed to create GitHub repo: %w", err)
	}

	return false, nil
}

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

func AddRemoteOrigin(repositoryName string, dir string) error {
	username := GetUsername()
	url := fmt.Sprintf("https://github.com/%s/%s.git", username, repositoryName)

	cmd := exec.Command("git", "remote", "add", "origin", url)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add remote: %w", err)
	}

	return nil
}

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
