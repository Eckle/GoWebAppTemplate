package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pbentes/80_20/src/db"
	"github.com/pbentes/80_20/src/routes"
)

func main() {
	db.InitDB()
	defer db.Cleanup()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.HTMLRender = &TemplRender{}

	router.Static("/assets", "./assets")

	routes.Set(router)

	address := "127.0.0.1:4321"
	err := router.Run(address)
	if err != nil {
		panic(err)
	}
}
