package controllers

import (
	model "halo_food/models"
	levelModel "halo_food/modules/level/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetLevelByID(db *gorm.DB, c echo.Context) error {
	idParam := c.Param("id_level")
	idLevel, err := strconv.Atoi(idParam)
	if err != nil {
		response := model.Response{
			ErrorCode: 1,
			Message:   "Parameter tidak boleh kosong !",
		}
		return c.JSON(http.StatusOK, response)
	}
	// levelModel.InitConnection(db)
	data := levelModel.GetLevelByID(db, idLevel)
	return c.JSON(http.StatusOK, data)
}

func GetAll(db *gorm.DB, c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	page, _ := strconv.Atoi(c.Param("page"))
	keywords := c.QueryParam("keywords")
	data := levelModel.GetAll(db, limit, page, keywords)

	return c.JSON(http.StatusOK, data)
}
