package users

import (
	"batch6/configs"
	"batch6/models"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	h := sha256.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	result := configs.DB.Create(&user)

	if result.Error != nil {
		BaseResponseError := models.BaseResponseError{
			Status:    false,
			Message:   "Failed insert to database",
			ErrorCode: 500,
		}

		return c.JSON(http.StatusInternalServerError, BaseResponseError)
	}

	successResponse := models.BaseResponseSuccess{
		Status:  true,
		Message: "Success",
		Data:    user,
	}

	return c.JSON(http.StatusCreated, successResponse)

}
