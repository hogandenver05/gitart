package repo

import (
	"fmt"
	"time"

)

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
