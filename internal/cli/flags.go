package cli

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Options struct {
	Message   string
	StartDate time.Time
	Target    int
	ArtPath   string
	SubCmd    string
}

func ParseFlags() *Options {
	options := &Options{}

	var startDateString string
	flag.StringVar(&options.Message, "m", "", "Message to display on GitHub contribution graph")
	flag.StringVar(&startDateString, "s", "", "Start date YYYY-MM-DD (top-left corner of the art)")
	flag.IntVar(&options.Target, "t", 0, "Target commits per day (shade of green)")
	flag.StringVar(&options.ArtPath, "p", "art", "Path to artwork repository")

	flag.Parse()

	if startDateString != "" {
		startDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			fmt.Println("Invalid start date format. Use YYYY-MM-DD. Exiting now")
			os.Exit(1)
		}
		options.StartDate = startDate
	}

	return options
}
