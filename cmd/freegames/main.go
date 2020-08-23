package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	freegames "github.com/arkiant/freegames/pkg"
	"github.com/arkiant/freegames/pkg/client/discord"
	"github.com/arkiant/freegames/pkg/platform/epicgames"
	"github.com/arkiant/freegames/pkg/service"
	"github.com/arkiant/freegames/pkg/storage/mongo"
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
	discordBot := discord.NewDiscordClient(&db, discordCommandHandler, dToken)

	// EXECUTE SERVICE
	fg := service.NewFreeGames(db)
	fg.AddPlatform(epicgames.NewEpicGames())
	fg.AddClient(discordBot)
	fg.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fg.Close()
}
