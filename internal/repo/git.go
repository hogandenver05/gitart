package repo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type NestedRepository struct {
	Path string
}

func NewNestedRepository(path string) *NestedRepository {
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		command := exec.Command("git", "init")
		command.Dir = path
		command.Run()
	}
	return &NestedRepository{Path: path}
}

func (repository *NestedRepository) CommitDay(day time.Time, count int) error {
	for i := 0; i < count; i++ {
		date := day.Format("2006-01-02T12:00:00")
		env := os.Environ()
		env = append(env, "GIT_AUTHOR_DATE="+date, "GIT_COMMITER_DATE="+date)
		command := exec.Command("git", "commit", "--allow-empty", "-m", fmt.Sprintf("pixel %s", date))
		command.Dir = repository.Path
		command.Env = env
		if err := command.Run(); err != nil {
			return err
		}
	}
	return nil
}
