package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	freegames "github.com/arkiant/freegames/internal"
	"github.com/arkiant/freegames/internal/getting"
	"github.com/arkiant/freegames/internal/platform/platform/epicgames"
	"github.com/arkiant/freegames/internal/platform/server/handler"
	"github.com/arkiant/freegames/internal/platform/storage/mongo"
	"github.com/arkiant/freegames/kit/cqrs/bus/inmemory"
	"github.com/arkiant/freegames/kit/http/server"
	"github.com/joho/godotenv"
)

// ENVIRONMENT VARIABLES
const (
	dataBaseURL  = "DATABASE_URL"
	discordToken = "DISCORD_TOKEN"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second
	dbTimeout       = 5 * time.Second
)

func Run() error {

	var (
		//commandBus = inmemory.NewCommandBus()
		eventBus = inmemory.NewEventBus()
		queryBus = inmemory.NewQueryBus()
	)

	var (
		_, base, _, _   = runtime.Caller(0)
		basePath        = filepath.Dir(base)
		environmentPath = filepath.Join(basePath, "../../../", ".env")
	)

	err := godotenv.Load(environmentPath)
	if err != nil {
		return err
	}

	dbURL := os.Getenv(dataBaseURL)
	if dbURL == "" {
		dbURL = "mongodb://localhost:27017"
	}

	freegamesRepository, err := mongo.NewMongoRepository(dbURL, "freegames", dbTimeout)
	if err != nil {
		return err
	}

	platforms := []freegames.Platform{
		epicgames.NewEpicGames(),
	}

	freegamesService := getting.NewFreegamesService(freegamesRepository, platforms, eventBus)
	freegamesQueryHandler := getting.NewFreegamesQueryHandler(freegamesService)
	queryBus.Register(getting.FregamesQueryType, freegamesQueryHandler)

	routes := []server.Route{
		server.NewRoute("GET", "freegames", handler.FreegamesHandler(queryBus)),
	}

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, routes)
	return srv.Run(ctx)
}
