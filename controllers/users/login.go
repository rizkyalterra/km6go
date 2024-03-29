package users

import (
	"batch6/middlewares"
	"batch6/models"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	h := sha256.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	// DB untuk cek login
	// pass login

	// generate Token JWT
	var userResponse models.UserLogin
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Id = user.Id
	userResponse.Token = middlewares.GenerateTokenJWT(user.Id, user.Name)

	successResponse := models.BaseResponseSuccess{
		Status:  true,
		Message: "Success",
		Data:    userResponse,
	}

	return c.JSON(http.StatusCreated, successResponse)

}
