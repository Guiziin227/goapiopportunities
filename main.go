package main

import (
	"fmt"
	"github.com/Guiziin227/goapiopportunities.git/config"
	"github.com/Guiziin227/goapiopportunities.git/router"
)

func main() {

	//Initialize config
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	//Inicialinzando o router
	router.Initialize()
}
