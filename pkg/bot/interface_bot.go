package bot

// Generic Bot Interface to implement bots for
// different platform under the same codebase
type Bot interface {
	// Run the bot
	Run() error
	// Stop the bot
	Stop() error
}
