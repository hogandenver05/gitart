package app

import "fmt"

type Grid [][]int

func BuildGrid(message string) (Grid, error) {
	const rows = 5
	grid := make(Grid, rows)
	for i := range rows {
		grid[i] = make([]int, 0)
	}

	font, err := LoadFont("internal/app/default-font")
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	for _, ch := range message {
		pattern, ok := font[ch]
		if !ok {
			continue
		}

		for r := range rows {
			grid[r] = append(grid[r], pattern[r]...)
			grid[r] = append(grid[r], 0)
		}
	}

	return grid, nil
}
