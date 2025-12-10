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

func (scheduler *Scheduler) Generate() error {
	rows := len(scheduler.Grid)
	if rows == 0 {
		return fmt.Errorf("grid has no rows")
	}

	cols := len(scheduler.Grid[0])

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
