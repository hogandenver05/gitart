package app

import "github.com/hogandenver05/gitart/internal/fonts"

type Grid [][]int

const gridRows = 5

// BuildGrid constructs a grid representation of the message using the predefined font map.
// Unsupported characters are skipped. Each rendered character is followed by a blank column.
func BuildGrid(message string) (Grid, error) {
	grid := newEmptyGrid(gridRows)

	for _, r := range message {
		pattern, exists := fonts.DefaultFontMap[r]
		if !exists {
			continue
		}

		appendPattern(grid, pattern)
	}

	return grid, nil
}

func newEmptyGrid(rows int) Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = []int{}
	}
	return grid
}

func appendPattern(grid Grid, pattern [][]int) {
	for rowIndex := range grid {
		if rowIndex < len(pattern) {
			grid[rowIndex] = append(grid[rowIndex], pattern[rowIndex]...)
		}
		grid[rowIndex] = append(grid[rowIndex], 0)
	}
}
