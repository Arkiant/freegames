package bootstrap

import (
	"context"
	"log"
	"time"

	main "github.com/arkiant/freegames"
	"github.com/arkiant/freegames/internal/getting/freegames"
	"github.com/arkiant/freegames/internal/platform/platform/epicgames"
	"github.com/arkiant/freegames/internal/platform/storage/mongo"
	"github.com/arkiant/freegames/kit/cqrs/bus/inmemory"
	"github.com/arkiant/freegames/kit/http/server"
	"github.com/joho/godotenv"
)

// ENVIRONMENT VARIABLES
const (
	dataBaseURL = "DATABASE_URL"
)

const (
	host            = "0.0.0.0"
	port            = 8080
	shutdownTimeout = 10 * time.Second
	dbTimeout       = 5 * time.Second
)

func Run() error {

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
		queryBus   = inmemory.NewQueryBus()
	)

	_ = eventBus

	file, err := main.ReadEnvironmentFile()
	if err != nil {
		panic(err)
	}

	environment, err := godotenv.Parse(file)
	if err != nil {
		return err
	}

	dbURL, ok := environment[dataBaseURL]
	if !ok {
		dbURL = "mongodb://localhost:27017"
		log.Println("Can't get the database url, getting the default value")
	}

	freegamesRepository, err := mongo.NewMongoRepository(dbURL, "freegames", dbTimeout)
	if err != nil {
		return err
	}

	freegamesService := freegames.NewService(freegamesRepository, epicgames.NewEpicGames())
	freegamesQueryHandler := freegames.NewQueryHandler(freegamesService)
	queryBus.Register(freegames.QueryType, freegamesQueryHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, Routes(queryBus, commandBus))
	return srv.Run(ctx)
}
