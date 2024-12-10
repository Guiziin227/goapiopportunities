package main

import (
	"github.com/Guiziin227/goapiopportunities.git/config"
	"github.com/Guiziin227/goapiopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	//Inicialinzando o router
	router.Initialize()
}
