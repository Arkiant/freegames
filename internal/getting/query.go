package getting

import (
	"context"
	"errors"

	"github.com/arkiant/freegames/kit/cqrs/query"
)

const FregamesQueryType query.Type = "query.getting.freegames"

type FreegamesQuery struct{}

func NewFreegamesQuery() FreegamesQuery {
	return FreegamesQuery{}
}

func (f FreegamesQuery) Type() query.Type {
	return FregamesQueryType
}

type FreegamesQueryHandler struct {
	service FreegamesService
}

func NewFreegamesQueryHandler(service FreegamesService) FreegamesQueryHandler {
	return FreegamesQueryHandler{service: service}
}

func (f FreegamesQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {

	_, ok := query.(FreegamesQuery)
	if !ok {
		return "", errors.New("unexpected query")
	}

	return f.service.GetFreeGames(ctx)

}
