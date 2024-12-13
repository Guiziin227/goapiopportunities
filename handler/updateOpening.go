package handler

import (
	"fmt"
	"github.com/Guiziin227/goapiopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.Bind(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"string").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}
	//Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary != 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	//Save opening
	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error updating opening: %s", err))
		return
	}
}
