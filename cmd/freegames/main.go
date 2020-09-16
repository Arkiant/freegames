package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/arkiant/freegames/pkg/client/discord"
	"github.com/arkiant/freegames/pkg/config/viper"
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

	var (
		_, base, _, _     = runtime.Caller(0)
		basePath          = filepath.Dir(base)
		environmentPath   = filepath.Join(basePath, "../../", ".env")
		configurationPath = filepath.Join(basePath, "../../", "config.yml")
	)

	err := godotenv.Load(environmentPath)
	if err != nil {
		log.Printf("Can't load .env file error: %s", err.Error())
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

	// Get configuration
	config, err := viper.NewViperConfiguration(configurationPath)

	// Create discord client
	discordBot := discord.NewDiscordClient(&db, dToken)

	// EXECUTE SERVICE
	fg := service.NewFreeGames(db, config)
	fg.AddPlatform(epicgames.NewEpicGames())
	fg.AddClient(discordBot)
	fg.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fg.Close()
}
