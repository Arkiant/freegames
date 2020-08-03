package main

import (
	"log"
	"os"
	"time"

	"github.com/arkiant/freegames/freegames"
	"github.com/arkiant/freegames/mongo"
	"github.com/arkiant/freegames/unreal"
	"github.com/joho/godotenv"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		if godotenv.Load(".env") != nil {
			panic("Can't load .env file")
		}

		dbURL = os.Getenv("DATABASE_URL")
	}

	db, err := mongo.NewMongoRepository(dbURL, "freegames", 5)
	if err != nil {
		panic(err)
	}

	const OnceADay = time.Hour * 24

	pool := []freegames.Platform{
		unreal.NewUnrealGames(),
	}

	// Once a day check for free games
	ticker := time.NewTicker(OnceADay)
	defer ticker.Stop()

	fg := freegames.GetAllFreeGames(pool, db)
	log.Printf("Found %v new free games", len(fg))

	for range ticker.C {
		fg = freegames.GetAllFreeGames(pool, db)
		log.Printf("Found %v new free games", len(fg))
	}

}
