package repo

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type NestedRepository struct {
	Path string
}

func NewNestedRepository(path string) (*NestedRepository, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	repository := &NestedRepository{Path: path}
	if err := repository.ResetRepository(); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	return repository, nil
}

func (repository *NestedRepository) CommitDay(day time.Time, count int) error {
	for range count {
		date := day.Format("2006-01-02T12:00:00")
		env := append(os.Environ(),
			"GIT_AUTHOR_DATE="+date,
			"GIT_COMMITTER_DATE="+date,
		)

		cmd := exec.Command("git", "commit", "--allow-empty", "-m", fmt.Sprintf("pixel %s", date))
		cmd.Dir = repository.Path
		cmd.Env = env
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to commit: %w", err)
		}
	}

	return nil
}
