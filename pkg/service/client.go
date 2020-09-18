package service

import (
	"log"

	freegames "github.com/arkiant/freegames/pkg"
)

// AddClient using chain patter we can add multiple clients to notify
func (f *Freegames) AddClient(client freegames.Client) *Freegames {
	f.clients = append(f.clients, client)
	return f
}

// executeClients Execute all clients suscribed
func executeClients(f *Freegames) {

	for _, v := range f.clients {
		err := v.Execute(f.config)
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
