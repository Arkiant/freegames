package main

import (
	"log"
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
		dataBaseURL  = "DATABASE_URL"
		discordToken = "DISCORD_TOKEN"
	)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Can't load .env file")
	}

	dbURL := os.Getenv(dataBaseURL)
	if dbURL == "" {
		dbURL = "mongodb://localhost:27017"
	}

	dToken := os.Getenv(discordToken)
	if dToken == "" {
		dbURL = "token"
	}

	db, err := mongo.NewMongoRepository(dbURL, "freegames", 5)
	if err != nil {
		panic(err)
	}

	// REGISTER CLIENT COMMANDS
	discordCommandHandler := freegames.NewCommandHandler()
	err = discordCommandHandler.Register("!freegames", discord.NewFreeGamesCommand())
	if err != nil {
		panic(err)
	}

	// TODO: ADD CLIENT FROM CONFIG FILE
	// TODO: BOT CONFIGURATION FROM CONFIG FILE
	discordBot := discord.NewDiscordClient(&db, discordCommandHandler).Configure(dToken)

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
