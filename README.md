<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/user-attachments/assets/7911011f-7111-4956-b6fd-31b39f032c52">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/user-attachments/assets/a2cf392a-c758-48bc-872e-8c19e04adae8">
  <img src="https://github.com/user-attachments/assets/a2cf392a-c758-48bc-872e-8c19e04adae8" alt="Gitart">
</picture>

Gitart is a command line tool that creates custom artwork on your GitHub contribution graph. You provide a message and the tool generates a series of commits inside a nested repository so the design appears on your profile.

---

## Features

• Create text-based artwork using a compact 3×5 pixel font  
• Choose a starting date and target number of commits for the darkest shade  
• Create a dedicated nested repository for artwork  
• Automatically generate empty commits that appear on your GitHub graph  

---

## Installation

### Debian / Ubuntu (recommended)

Download the `.deb` from the GitHub Releases page:

https://github.com/hogandenver05/gitart/releases

Install it:

```bash
sudo dpkg -i gitart_<version>_amd64.deb
sudo apt -f install
```

Run:

```bash
gitart
```

---

### From source (development)

Clone the repository:

```bash
git clone https://github.com/hogandenver05/gitart
cd gitart
```

Build:

```bash
go build -o gitart ./cmd/gitart
```

Run locally:

```bash
./gitart
```

Optional global install:

```bash
sudo mv gitart /usr/local/bin
gitart
```

---

## Usage

### Basic usage

Run interactively (prompts for missing options):

```bash
gitart
```

Provide all options via flags:

```bash
gitart -m "HELLO" -s 2025-01-01 -t 1 --push
```

---

### Command-line flags

- `-m, --message`  
  Message to display on the GitHub contribution graph

- `-s, --start-date`  
  Start date in `YYYY-MM-DD` format (top-left corner of the artwork)

- `-t, --target`  
  Target number of commits per day (higher = darker shade)

- `-p, --path`  
  Path to the artwork repository (default: `art`)

- `--push`  
  Automatically push commits to GitHub

- `--private`  
  Used with `--push` to create a private repository

- `--no-reset`  
  Used with `--push` to preserve existing commits and layer messages

- `--no-count`  
  Disable contribution counting and always commit the exact target count

---

### Examples

Generate artwork without pushing:

```bash
gitart -m "FOX SPORTS" -s 2024-12-16 -t 1
```

Generate and push to GitHub:

```bash
gitart -m "HELLO WORLD" -s 2024-01-01 -t 1 --push
```

Create a private repository:

```bash
gitart -m "HELLO WORLD" -s 2024-01-01 -t 1 --push --private
```

Layer messages for special effects:

```bash
gitart -m "SPECIAL" -s 2024-01-01 -t 1 --push
gitart -m "SPECIAL" -s 2024-01-08 -t 1 --push --no-reset
```

Disable contribution counting:

```bash
gitart -m "HELLO" -s 2024-01-01 -t 5 --push --no-count
```

---

## How it works

Gitart renders your message using a small 3×5 font.  
Each pixel represents a day on the GitHub contribution graph.

Filled pixels generate commits.  
Blank pixels remain empty.

The tool creates a nested Git repository dedicated to the artwork.  
Once commits are generated, pushing the repository causes the artwork to appear on your GitHub profile.

By default, gitart checks your existing contributions for each day and adjusts the number of new commits so the total does not exceed the target shade.  
Use `--no-count` to disable this behavior.

---

## Requirements

• Go 1.21 or later (for building from source)  
• Git  

---

## Contributing

Pull requests are welcome.  
Please follow the existing project structure and keep Go source separate from packaging artifacts.

---

## License

Distributed under the MIT License.  
See the LICENSE file for details.
