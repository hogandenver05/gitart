package repo

import (
	"fmt"
	"time"
)

// GenerateInstructions returns formatted instructions for manually pushing the repository to GitHub.
func GenerateInstructions() string {
	username := GetUsername()
	date := time.Now().Format("2006-01-02")
	
	return fmt.Sprintf(`
Your contribution art has been generated in ./art/

To show it on GitHub:

1. Create a new GitHub repository
2. Set the remote for the local repository
3. Push the commits

   gh repo create %s/gitart-%s --public
   cd art
   git remote add origin https://github.com/%s/gitart-%s
   git branch -M main
   git push -u origin main

Your contribution graph will update automatically after pushing!
`, username, date, username, date)
}

// FormatPushStatus returns a formatted string describing the result of a GitHub push operation.
func FormatPushStatus(status *PushStatus) string {
	note := ""
	if status.RepoAlreadyExists {
		note = "\n note: repository already existed on GitHub"
	}
	
	return fmt.Sprintf(`
GitHub push complete
 repository name: %s
 local path: %s
 github user: %s
 remote url: %s
 branch: %s%s
 Check it out at https://github.com/%s
`, status.RepositoryName, status.RepositoryPath, status.Username, status.RemoteURL, status.Branch, note, status.Username)
}
