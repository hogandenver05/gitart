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

	repository, err := repo.NewNestedRepository(options.ArtPath)
	if err != nil {
		printErrorAndExit(err)
	}

	scheduler := app.NewScheduler(grid, options.StartDate, options.Target, repository)
	if err := scheduler.Generate(); err != nil {
		printErrorAndExit(err)
	}

	if options.Push {
		status, err := repository.PushToGitHub(options.Private, !options.NoReset)
		if err != nil {
			printErrorAndExit(err)
		}

		fmt.Println("GitHub push complete")
		fmt.Println(" repository name:", status.RepositoryName)
		fmt.Println(" local path:", status.RepositoryPath)
		fmt.Println(" github user:", status.Username)
		fmt.Println(" remote url:", status.RemoteURL)
		fmt.Println(" branch:", status.Branch)

		if status.RepoAlreadyExists {
			fmt.Println(" note: repository already existed on GitHub")
		}

		fmt.Println(" Check it out at https://github.com/" + status.Username)
		return
	}
}

func printErrorAndExit(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
