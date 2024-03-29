package main

import (
	"batch6/configs"
	"batch6/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8080")
}
