package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa um router com as configuracoes default do gin
	//Initialize router
	router := gin.Default()
	//Define a rota
	//Initialize routes
	initializeRoutes(router)
	//run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
