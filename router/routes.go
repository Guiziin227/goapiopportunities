package router

import (
	"github.com/Guiziin227/goapiopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize a router group
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		//Show opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}
}
