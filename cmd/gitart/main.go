package main

import (
	"fmt"
	"os"

	"github.com/hogandenver05/gitart/internal/app"
	"github.com/hogandenver05/gitart/internal/cli"
	"github.com/hogandenver05/gitart/internal/repo"
)

func main() {
	options, err := cli.ParseFlagsOrPrompt()
	if err != nil {
		printErrorAndExit(err)
	}

	grid, err := app.BuildGrid(options.Message)
	if err != nil {
		printErrorAndExit(err)
	}

	repository, err := repo.NewNestedRepository(options.ArtPath, options.NoCount)
	if err != nil {
		printErrorAndExit(err)
	}

	scheduler := app.NewScheduler(grid, options.StartDate, options.Target, repository)
	if err := scheduler.Generate(); err != nil {
		printErrorAndExit(err)
	}

	if err := repository.IncludeREADMEIfPresent(); err != nil {
		printErrorAndExit(err)
	}

	if options.Push {
		regenerateArtwork := func() error {
			return scheduler.Generate()
		}
		status, err := repository.PushToGitHub(options.Private, !options.NoReset, regenerateArtwork)
		if err != nil {
			printErrorAndExit(err)
		}

		fmt.Println(repo.FormatPushStatus(status))
		return
	} else {
		fmt.Println(repo.GenerateInstructions())
	}
}

func printErrorAndExit(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
