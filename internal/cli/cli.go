package cli

import (
	"fmt"
	"time"
)

type Options struct {
	Message   string
	StartDate time.Time
	Target    int
	ArtPath   string
	Push      bool
	Private   bool
	NoReset   bool
	NoCount   bool
}

// ParseFlagsOrPrompt parses command-line flags and prompts for any missing required options.
func ParseFlagsOrPrompt() (*Options, error) {
	options, err := ParseFlags()
	if err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}
	
	return PromptOptions(options)
}
