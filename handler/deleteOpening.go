package handler

import (
	"fmt"
	"github.com/Guiziin227/goapiopportunities.git/schemas"
	"github.com/gin-gonic/gin"

	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"string").Error())
		return
	}

	opening := schemas.Opening{}

	//Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}
	//Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening: %s", err))
		return
	}
	sendSuccess(ctx, "delete", opening)
}
