package cli

import (
	"flag"
	"fmt"
	"time"
)

// ParseFlags parses command-line flags and returns an Options struct.
// Dates are normalized to UTC at noon to ensure consistent representation.
func ParseFlags() (*Options, error) {
	options := &Options{}

	var startDateString string
	flag.StringVar(&options.Message, "m", "", "Message to display on GitHub contribution graph")
	flag.StringVar(&startDateString, "s", "", "Start date YYYY-MM-DD (top-left corner of the art)")
	flag.IntVar(&options.Target, "t", 0, "Target commits per day (shade of green)")
	flag.StringVar(&options.ArtPath, "p", "art", "Path to artwork repository")
	flag.BoolVar(&options.Push, "push", false, "Automatically push commits to GitHub")
	flag.BoolVar(&options.Private, "private", false, "Used with --push to set repository visibility to private")
	flag.BoolVar(&options.NoReset, "no-reset", false, "Used with --push to avoid resetting the repository")
	flag.BoolVar(&options.Count, "count", false, "Enable contribution counting (limit commits to target count by counting existing contributions)")

	flag.Parse()

	if startDateString != "" {
		parsedDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format, use YYYY-MM-DD")
		}
		
		year, month, day := parsedDate.UTC().Date()
		options.StartDate = time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	}

	return options, nil
}
