package freegames

// Context freegames specific, in this context we save current channel and arguments passed to commands
type Context struct {
	Channel string
	Args    []string
}
