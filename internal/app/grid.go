package app

type Grid [][]int

var Font = map[rune]Grid{
	'H': {
		{1, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
		{1, 0, 1},
	},
	'I': {
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	},
	'R': {
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
		{1, 0, 1},
		{1, 0, 1},
	},
	'E': {
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
	},
	'M': {
		{1, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
		{1, 0, 1},
		{1, 0, 1},
	},
	' ': {
		{0},
		{0},
		{0},
		{0},
		{0},
	},
}

func BuildGrid(message string) Grid {
	rows := 5
	grid := make(Grid, rows)

	for _, character := range message {
		pattern, ok := Font[character]

		if !ok {
			pattern = Font[' ']
		}

		for r := range rows {
			grid[r] = append(grid[r], pattern[r]...)
			grid[r] = append(grid[r], 0)
		}
	}

	return grid
}
