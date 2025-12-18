package repo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type NestedRepository struct {
	Path           string
	EnableCounting bool
}

// NewNestedRepository creates a new repository at the specified path.
// Only initializes git if the repository doesn't exist, preserving existing commits when using --no-reset.
func NewNestedRepository(path string, enableCounting bool) (*NestedRepository, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	repository := &NestedRepository{
		Path:           path,
		EnableCounting: enableCounting,
	}

	gitPath := filepath.Join(path, ".git")
	if _, err := os.Stat(gitPath); os.IsNotExist(err) {
		if err := repository.ResetRepository(); err != nil {
			return nil, fmt.Errorf("failed to initialize repository: %w", err)
		}
	}

	return repository, nil
}

// CommitDay creates the specified number of empty commits for the given day.
// Dates are normalized to UTC at noon to ensure consistent representation on GitHub's contribution graph.
// The commit count is adjusted based on existing contributions to avoid exceeding the target, unless counting is disabled.
func (repository *NestedRepository) CommitDay(day time.Time, count int) error {
	commitsToMake := count

	if repository.EnableCounting {
		existingContributions, err := GetContributionCount(day)
		if err != nil {
			return fmt.Errorf("failed to get contribution count for %s: %v", day.Format("2006-01-02"), err)
		}

		commitsToMake = max(count - existingContributions, 0)

		if commitsToMake == 0 {
			return nil
		}
	}

	utcDay := day.UTC()
	year, month, date := utcDay.Date()
	normalizedDay := time.Date(year, month, date, 12, 0, 0, 0, time.UTC)

	for range commitsToMake {
		dateString := normalizedDay.Format("2006-01-02T15:04:05Z")
		env := append(os.Environ(),
			"GIT_AUTHOR_DATE="+dateString,
			"GIT_COMMITTER_DATE="+dateString,
		)

		commitMessageDateString := normalizedDay.Format("2006-01-02")
		cmd := exec.Command("git", "commit", "--allow-empty", "-m", fmt.Sprintf("pixel %s", commitMessageDateString))
		cmd.Dir = repository.Path
		cmd.Env = env
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to commit: %w", err)
		}
	}

	return nil
}

// IncludeREADMEIfPresent adds README.md to the repository if it exists in the repository path.
// The README is committed with the current date to avoid interfering with artwork commits
// that use historical dates for the contribution graph pattern.
func (repository *NestedRepository) IncludeREADMEIfPresent() error {
	readmePath := filepath.Join(repository.Path, "README.md")

	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		return nil
	}

	addCommand := exec.Command("git", "add", "README.md")
	addCommand.Dir = repository.Path
	if err := addCommand.Run(); err != nil {
		return fmt.Errorf("failed to stage README.md: %w", err)
	}

	checkCommand := exec.Command("git", "diff", "--cached", "--exit-code", "README.md")
	checkCommand.Dir = repository.Path
	if err := checkCommand.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) && exitError.ExitCode() == 1 {
		} else {
			return fmt.Errorf("failed to check README.md status: %w", err)
		}
	} else {
		return nil
	}

	commitCommand := exec.Command("git", "commit", "-m", "Add README.md")
	commitCommand.Dir = repository.Path
	if err := commitCommand.Run(); err != nil {
		return fmt.Errorf("failed to commit README.md: %w", err)
	}

	return nil
}
