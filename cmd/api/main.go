package main

import (
	"log"

	"github.com/arkiant/freegames/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
