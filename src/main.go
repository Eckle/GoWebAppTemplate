package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pbentes/Chime/src/db"
	"github.com/pbentes/Chime/src/router"
)

func main() {
	db.InitDB()
	defer db.Cleanup()

	e := echo.New()

	e.Static("/", "/assets")
	router.Set(e)

	e.Logger.Fatal(e.Start("127.0.0.1:4321"))
}
