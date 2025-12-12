package app

import (
	"fmt"
	"time"

	"github.com/hogandenver05/gitart/internal/repo"
)

type Scheduler struct {
	Grid       Grid
	StartDate  time.Time
	Target     int
	Repository *repo.NestedRepository
}

// NewScheduler creates a new scheduler for generating artwork commits.
func NewScheduler(
	grid Grid,
	startDate time.Time,
	target int,
	repository *repo.NestedRepository,
) *Scheduler {
	return &Scheduler{
		Grid:       grid,
		StartDate:  startDate,
		Target:     target,
		Repository: repository,
	}
}

// Generate creates commits for all grid positions marked with 1, mapping them to dates on the contribution graph.
// Each column represents a week (7 days), and each row represents a day offset within that week.
func (scheduler *Scheduler) Generate() error {
	rows := len(scheduler.Grid)
	if rows == 0 {
		return fmt.Errorf("grid has no rows")
	}

	cols := len(scheduler.Grid[0])

	for i, row := range scheduler.Grid {
		if len(row) != cols {
			return fmt.Errorf("grid row %d has length %d, expected %d", i, len(row), cols)
		}
	}

	for col := range cols {
		for row := range rows {
			if scheduler.Grid[row][col] == 1 {
				day := scheduler.StartDate.AddDate(0, 0, col*7+row)
				err := scheduler.Repository.CommitDay(day, scheduler.Target)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
