package cli

import (
	"flag"
	"fmt"
	"time"
)

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

	flag.Parse()

	if startDateString != "" {
		startDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format, use YYYY-MM-DD")
		}
		
		options.StartDate = startDate
	}

	return options, nil
}
