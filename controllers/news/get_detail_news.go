package news

import (
	"batch6/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDetailNewsController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	errResponse := models.BaseResponseError{}
	if err != nil {
		errResponse.ErrorCode = 4031
		errResponse.Message = "Error in parameter url"
		errResponse.Status = false
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	detailNews := models.News{
		Id:      idInt,
		Title:   "Detail News",
		Content: "Lorem Ipsum Dolor Sit Amet",
	}

	successResponse := models.BaseResponseSuccess{
		Message: "Success",
		Status:  true,
		Data:    detailNews,
	}

	return c.JSON(http.StatusOK, successResponse)
}
