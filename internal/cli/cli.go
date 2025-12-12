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
	SubCmd    string
	Push      bool
	Private   bool
	NoReset   bool
}

func ParseFlagsOrPrompt() (*Options, error) {
	options, err := ParseFlags()
	if err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}
	
	return PromptOptions(options)
}
