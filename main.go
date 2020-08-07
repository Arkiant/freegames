package main

import (
	"os"

	"github.com/arkiant/freegames/epicgames"
	"github.com/arkiant/freegames/freegames"
	"github.com/arkiant/freegames/mongo"
	"github.com/joho/godotenv"
)

func main() {

	// ENVIRONMENT VARIABLES
	const (
		dataBaseURL = "DATABASE_URL"
	)

	dbURL := os.Getenv(dataBaseURL)
	if dbURL == "" {
		if godotenv.Load(".env") != nil {
			panic("Can't load .env file")
		}

		dbURL = os.Getenv(dataBaseURL)
		if dbURL == "" {
			dbURL = "mongodb://localhost:27017"
		}
	}

	db, err := mongo.NewMongoRepository(dbURL, "freegames", 5)
	if err != nil {
		panic(err)
	}

	fg := freegames.NewFreeGames(&db)
	fg.AddPlatform(epicgames.NewEpicGames()).Run()

}
