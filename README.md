<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/user-attachments/assets/7911011f-7111-4956-b6fd-31b39f032c52">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/user-attachments/assets/a2cf392a-c758-48bc-872e-8c19e04adae8">
  <img src="https://github.com/user-attachments/assets/a2cf392a-c758-48bc-872e-8c19e04adae8" alt="Gitart">
</picture>

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
go build -o gitart ./cmd/gitart/main.go
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

## Usage

### Basic Usage

Run gitart interactively (will prompt for missing options):

```
./gitart
```

Or provide all options via command-line flags:

```
./gitart -m "HELLO" -s 2025-01-01 -t 1 --push
```

### Command-Line Flags

- `-m, --message`: Message to display on GitHub contribution graph
- `-s, --start-date`: Start date in YYYY-MM-DD format (top-left corner of the artwork)
- `-t, --target`: Target number of commits per day (higher = darker shade of green)
- `-p, --path`: Path to artwork repository (default: `art`)
- `--push`: Automatically push commits to GitHub
- `--private`: Used with `--push` to create a private repository
- `--no-reset`: Used with `--push` to preserve existing commits and layer messages
- `--no-count`: Disable contribution counting (use target count regardless of existing contributions)

### Examples

Generate artwork without pushing:

```
./gitart -m "FOX SPORTS" -s 2024-12-16 -t 1
```

Generate and automatically push to GitHub:

```
./gitart -m "HELLO WORLD" -s 2024-01-01 -t 1 --push
```

Create a private repository:

```
./gitart -m "HELLO WORLD" -s 2024-01-01 -t 1 --push --private
```

Layer multiple messages for special effects (preserve existing commits):

```
./gitart -m "SPECIAL" -s 2024-01-01 -t 1 --push
./gitart -m "SPECIAL" -s 2024-01-08 -t 1 --push --no-reset
```

Disable contribution counting (use exact target count):

```
./gitart -m "HELLO" -s 2024-01-01 -t 5 --push --no-count
```

## How it works

Gitart renders your message using a small three by five font.
Each pixel represents a day on the contribution graph.
Filled pixels produce commits.
Blank pixels remain empty.

The tool creates a nested Git repository dedicated to the artwork.
Once commits are generated, pushing it to GitHub displays the artwork on your profile.

By default, gitart automatically checks your existing GitHub contributions for each day and adjusts the commit count to ensure the total (existing + new) doesn't exceed your target. This prevents artwork commits from pushing days beyond the desired shade level. Use the `--no-count` flag to disable this behavior and commit the exact target count regardless of existing contributions.


## Requirements

• Go version 1.21 or later  
• Git  

## Contributing

Pull requests are welcome. Please follow the existing project structure.

## License

Distributed under the MIT License.  
See the LICENSE file for details.
