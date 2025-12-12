package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// PromptOptions prompts the user for any missing required options.
// Dates are normalized to UTC at noon to ensure consistent representation.
func PromptOptions(options *Options) (*Options, error) {
	reader := bufio.NewReader(os.Stdin)

	if options.Message == "" {
		fmt.Print("Enter your message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read message: %w", err)
		}
		if message == "" {
			return nil, fmt.Errorf("message cannot be empty")
		}
		options.Message = strings.TrimSpace(message)
	}

	if options.StartDate.IsZero() {
		fmt.Print("Enter start date YYYY-MM-DD (top-left corner of the artwork): ")
		startDateString, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read start date: %w", err)
		}
		startDateString = strings.TrimSpace(startDateString)
		parsedDate, err := time.Parse("2006-01-02", startDateString)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format, use YYYY-MM-DD")
		}

		year, month, day := parsedDate.UTC().Date()
		options.StartDate = time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	}

	if options.Target <= 0 {
		fmt.Print("Enter target commits per day (more = darker shade): ")
		targetString, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read target: %w", err)
		}
		targetString = strings.TrimSpace(targetString)
		target, err := strconv.Atoi(targetString)
		if err != nil || target <= 0 {
			return nil, fmt.Errorf("invalid number, must be greater than 0")
		}
		options.Target = target
	}

	if options.ArtPath == "" {
		options.ArtPath = "art"
	}

	fmt.Println("Working...")

	return options, nil
}
