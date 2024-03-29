package routes

import (
	"batch6/constants"
	"batch6/controllers/news"
	"batch6/controllers/users"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())

	eAuth := e.Group("")
	e.Use(echojwt.JWT([]byte(constants.PRIVATE_KEY_JWT)))
	eAuth.GET("/news", news.GetNewsController)
	eAuth.GET("/news/:id", news.GetDetailNewsController)
	eAuth.POST("/news", news.AddNewsController)

	e.POST("/register", users.RegisterController)
	e.POST("/login", users.LoginController)

}
