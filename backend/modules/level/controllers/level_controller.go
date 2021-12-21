package controllers

import (
	model "halo_food/model"
	levelModel "halo_food/modules/level/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetLevelByID(c echo.Context) error {
	idParam := c.Param("id_level")
	idLevel, err := strconv.Atoi(idParam)
	if err != nil {
		response := model.Response{
			ErrorCode: 1,
			Message:   "Parameter tidak boleh kosong !",
		}
		return c.JSON(http.StatusOK, response)
	}
	// levelModel.InitConnection(config.DB)
	data := levelModel.GetLevelByID(idLevel)
	return c.JSON(http.StatusOK, data)
}

func GetAll(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	page, _ := strconv.Atoi(c.Param("page"))
	keywords := c.QueryParam("keywords")
	// levelModel.InitConnection(config.DB)
	data := levelModel.GetAll(limit, page, keywords)

	return c.JSON(http.StatusOK, data)
}
