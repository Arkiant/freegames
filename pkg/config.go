package freegames

// Configuration config structure
type Configuration struct {
	Clients  map[string]ClientConfig
	Platform []string
}

// Config abstraction to decouple and create an adapter to use any configuration library
type Config interface {
	ClientConfig() (map[string]ClientConfig, error)
}
