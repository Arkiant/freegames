package freegames

import (
	"context"
	"errors"

	"github.com/arkiant/freegames/kit/cqrs/query"
)

const QueryType query.Type = "query.getting.freegames"

type Query struct{}

func NewQuery() Query {
	return Query{}
}

func (f Query) Type() query.Type {
	return QueryType
}

type QueryHandler struct {
	service Service
}

func NewQueryHandler(service Service) QueryHandler {
	return QueryHandler{service: service}
}

func (f QueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {

	_, ok := query.(Query)
	if !ok {
		return "", errors.New("unexpected query")
	}

	return f.service.GetFreeGames(ctx)

}
