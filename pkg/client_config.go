package freegames

// ClientConfig domain representation
type ClientConfig interface {
	GetChannel() string
	GetToken() string
	GetPrefix() string
}
