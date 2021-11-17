package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/arkiant/freegames/internal/platform/bus/inmemory"
	"github.com/arkiant/freegames/internal/platform/server"
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
		commandBus = inmemory.NewCommandBus()
		// eventBus   = inmemory.NewEventBus()
	)

	var (
		_, base, _, _   = runtime.Caller(0)
		basePath        = filepath.Dir(base)
		environmentPath = filepath.Join(basePath, "../../", ".env")
	)

	err := godotenv.Load(environmentPath)
	if err != nil {
		return err
	}

	dbURL := os.Getenv(dataBaseURL)
	if dbURL == "" {
		dbURL = "mongodb://localhost:27017"
	}

	dToken := os.Getenv(discordToken)
	if dToken == "" {
		dbURL = "token"
	}

	// db, err := mongo.NewMongoRepository(dbURL, "freegames", 5)
	// if err != nil {
	// 	return err
	// }

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
