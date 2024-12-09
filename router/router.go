package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa um router com as configuracoes default do gin
	router := gin.Default()
	//Define a rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//run
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
