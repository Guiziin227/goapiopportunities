package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.Bind(&request)
	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating a new opening request: %v", err.Error())
		return
	}
}
