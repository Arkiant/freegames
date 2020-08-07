package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/arkiant/freegames/discord"
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

	// TODO: ADD CLIENT FROM CONFIG FILE
	// TODO: BOT CONFIGURATION FROM CONFIG FILE
	discordBot := discord.NewDiscordClient(&db).Configure("token")

	// EXECUTE SERVICE
	fg := freegames.NewFreeGames(&db)
	fg.AddPlatform(epicgames.NewEpicGames())
	fg.AddClient(discordBot)
	fg.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fg.Close()
}
