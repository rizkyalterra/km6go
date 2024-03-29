package news

import (
	"batch6/configs"
	"batch6/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNewsController(c echo.Context) error {
	news := []models.News{}

	result := configs.DB.Find(&news)

	if result.Error != nil {
		BaseResponseError := models.BaseResponseError{
			Status:    false,
			Message:   "Failed get news from database",
			ErrorCode: 500,
		}

		return c.JSON(http.StatusInternalServerError, BaseResponseError)
	}

	successResponse := models.BaseResponseSuccess{
		Status:  true,
		Message: "Success",
		Data:    news,
	}

	return c.JSON(http.StatusCreated, successResponse)
}
