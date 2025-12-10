package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func PromptOptions(options *Options) *Options {
	reader := bufio.NewReader(os.Stdin)

	if options.Message == "" {
		fmt.Print("Enter your message: ")
		message, _ := reader.ReadString('\n')
		if message == "" {
			fmt.Println("Message cannot be empty. Exiting now.")
			os.Exit(1)
		}
		options.Message = strings.TrimSpace(message)
	}

	if options.StartDate.IsZero() {
		fmt.Print("Enter start date YYYY-MM-DD (top-left corner of the artwork): ")
		startDateString, _ := reader.ReadString('\n')
		startDateString = strings.TrimSpace(startDateString)
		startDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			fmt.Println("Invalid start date format. Use YYYY-MM-DD. Exiting now.")
			os.Exit(1)
		}
		options.StartDate = startDate
	}

	if options.Target <= 0 {
		fmt.Print("Enter target commits per day (more = darker shade): ")
		targetString, _ := reader.ReadString('\n')
		targetString = strings.TrimSpace(targetString)
		target, err := strconv.Atoi(targetString)
		if err != nil || target <= 0 {
			fmt.Println("Invalid number. Must be greater than 0. Exiting now.")
			os.Exit(1)
		}
		options.Target = target
	}

	if options.ArtPath == "" {
		options.ArtPath = "art"
	}

	return options
}
