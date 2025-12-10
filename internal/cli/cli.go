package cli

func ParseFlagsOrPrompt() *Options {
	options := ParseFlags()
	return PromptOptions(options)
}
