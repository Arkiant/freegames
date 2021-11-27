package freegames

import (
	"net/http"

	"github.com/arkiant/freegames/internal/getting/freegames"
	"github.com/arkiant/freegames/kit/cqrs/query"
	"github.com/gin-gonic/gin"
)

func Handler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := queryBus.Dispatch(ctx, freegames.NewQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
