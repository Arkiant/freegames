package freegames

// Configuration config structure
type Configuration struct {
	Clients []struct {
		Discord struct {
			Enable  bool
			Channel string
		}
	}
	Platform []string
}

// Config abstraction to decouple and create an adapter to use any configuration library
type Config interface {
	ClientConfig() ([]struct {
		Discord struct {
			Enable  bool
			Channel string
		}
	}, error)
}
