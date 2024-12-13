package handler

import (
	"fmt"
	"github.com/Guiziin227/goapiopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error listing opening: %s", err))
		return
	}
	sendSuccess(ctx, "list", opening)
}
