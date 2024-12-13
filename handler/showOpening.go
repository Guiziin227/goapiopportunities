package handler

import (
	"fmt"
	"github.com/Guiziin227/goapiopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"string").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
	}

	sendSuccess(ctx, "show", opening)
}
