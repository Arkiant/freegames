package main

import (
	"os"

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

	fg := freegames.NewFreeGames(&db)
	fg.AddPlatform(unreal.NewUnrealGames())
	fg.Run()

}
