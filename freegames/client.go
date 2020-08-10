package freegames

import (
	"log"
)

// Client abstraction
type Client interface {
	Execute() error
	GetName() string
	Close()
	SendMessage() error
	SendMessageToChannel(string) error
}

// AddClient using chain patter we can add multiple clients to notify
func (f *Freegames) AddClient(client Client) *Freegames {
	f.clients = append(f.clients, client)
	return f
}

// executeClients Execute all clients suscribed
func executeClients(f *Freegames) {

	for _, v := range f.clients {
		err := v.Execute()
		if err != nil {
			log.Fatalf("Error when try to execute %s bot, error %s", v.GetName(), err.Error())
		}
		log.Printf("Executed: %s bot", v.GetName())
	}
}

func closeClients(f *Freegames) {
	for _, v := range f.clients {
		v.Close()
	}
}
