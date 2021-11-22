package handler

import (
	"net/http"

	"github.com/arkiant/freegames/internal/getting"
	"github.com/arkiant/freegames/kit/cqrs/query"
	"github.com/gin-gonic/gin"
)

func FreegamesHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := queryBus.Dispatch(ctx, getting.NewFreegamesQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
