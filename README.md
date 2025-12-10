
<img width="1404" alt="gitart" src="https://github.com/user-attachments/assets/b5297f1b-6395-48c5-bc80-ce6008a5157a" />

Gitart is a command line tool that creates custom artwork on your GitHub contribution graph. You provide a message and the tool generates a series of commits inside a nested repository so the design appears on your profile.

## Features

• Create text based artwork with a compact three by five pixel font  
• Choose your starting date and target number of commits for darkest color  
• Create a nested repository for artwork  
• Automatically generate empty commits that show up on your GitHub graph  

## Installation

Clone the repository

```
git clone https://github.com/hogandenver05/gitart
cd gitart
```

Build

```
go build ./cmd/gitart
```

Global install and run
```
sudo mv gitart /usr/local/bin
gitart
```

Or run

```
./gitart
```


## How it works

Gitart renders your message using a small three by five font.
Each pixel represents a day on the contribution graph.
Filled pixels produce commits.
Blank pixels remain empty.

The tool creates a nested Git repository dedicated to the artwork.
Once commits are generated, pushing it to GitHub displays the artwork on your profile.


## Requirements

• Go version 1.21 or later  
• Git  

## Contributing

Pull requests are welcome. Please follow the existing project structure.

## License

Distributed under the MIT License.  
See the LICENSE file for details.
