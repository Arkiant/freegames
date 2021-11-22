package bootstrap

import (
	"github.com/arkiant/freegames/internal/platform/server/handler"
	"github.com/arkiant/freegames/kit/cqrs/command"
	"github.com/arkiant/freegames/kit/cqrs/query"
	"github.com/arkiant/freegames/kit/http/server"
)

func Routes(queryBus query.Bus, commandBus command.Bus) []server.Route {
	return []server.Route{
		server.NewRoute("GET", "freegames", handler.FreegamesHandler(queryBus)),
	}
}
