package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hogandenver05/gitart/internal/app"
	"github.com/hogandenver05/gitart/internal/cli"
	"github.com/hogandenver05/gitart/internal/repo"
)

func main() {
	options := cli.ParseFlagsOrPrompt()
	grid := app.BuildGrid(options.Message)
	artRepo := repo.NewNestedRepository(options.ArtPath)
	scheduler := app.NewScheduler(grid, options.StartDate, options.Target, artRepo)

	if err := scheduler.Generate(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println(`
	Your contribution art has been generated in ./art/

	To show it on GitHub:

	1. Create a new GitHub repository
	2. Set the remote for the local repository
	3. Push the commits
	   
	   gh repo create <your_username>/gitart-` + time.Now().Format("2006-01-02") + ` --public
	   cd art
	   git remote add origin <github_repo_url>
	   git branch -M main
	   git push -u origin main
	
	Your contribution graph will update automatically after pushing!
	`)

}
