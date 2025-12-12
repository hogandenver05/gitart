package app

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func LoadFont(path string) (map[rune]Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open font file: %w", err)
	}
	defer file.Close()

	out := make(map[rune]Grid)
	scanner := bufio.NewScanner(file)

	var current rune
	var buffer []string

	commit := func() {
		if current == 0 {
			return
		}
		grid := make(Grid, len(buffer))
		for i, row := range buffer {
			grid[i] = make([]int, len(row))
			for j, ch := range row {
				if ch == '1' {
					grid[i][j] = 1
				}
			}
		}
		out[current] = grid
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			commit()
			current = 0
			buffer = buffer[:0]
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			commit()
			r, _ := utf8.DecodeRuneInString(line[1 : len(line)-1])
			current = r
			buffer = buffer[:0]
			continue
		}

		buffer = append(buffer, line)
	}

	commit()

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("cannot read font file: %w", err)
	}

	return out, nil
}
