package controllers

import (
	"halo_food/config"
	general "halo_food/helpers/general"
	"halo_food/model"
	usermodel "halo_food/modules/users/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	UserValidator struct {
		Username string `json:"username" form:"username" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}
)

func DoLogin(c echo.Context) error {
	validate := validator.New()
	userValidator := new(UserValidator)
	if err := c.Bind(userValidator); err != nil {
		return err
	}

	err := validate.Struct(userValidator)
	if err != nil {
		var errField []model.ErrorValidator
		for _, e := range err.(validator.ValidationErrors) {
			errField = append(errField, model.ErrorValidator{
				FieldName:    e.Field(),
				ErrorMessage: general.ErrorMessage(e.Field(), e.Tag()),
			})
			response := model.Response{
				ErrorCode: 101,
				Message:   "Gagal login, check field yang wajib terisi",
				Errors:    errField,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	if usermodel.DoLoginUser(config.DB, userValidator.Username, userValidator.Password) {
		user, _ := usermodel.GetOneByUsername(config.DB, userValidator.Username)
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusOK, false)
}

// func GetAll(c echo.Context) error {
// 	limit, _ := strconv.Atoi(c.Param("limit"))
// 	page, _ := strconv.Atoi(c.Param("page"))
// 	keywords := c.QueryParam("keywords")
// 	data := filmModel.GetAll(config.DB, limit, page, keywords)

// 	return c.JSON(http.StatusOK, data)
// }

// func Test(){
// 	validate := validator.New()
// 	var user model.Users
// 	userValidator := new(UserValidator)

// 	general.BindData(c, &user)
// 	if err := c.Bind(userValidator); err != nil {
// 		return err
// 	}

// 	err := validate.Struct(userValidator)
// 	if err != nil {
// 		var errField []model.ErrorValidator
// 		for _, e := range err.(validator.ValidationErrors) {
// 			errField = append(errField, model.ErrorValidator{
// 				FieldName:    e.Field(),
// 				ErrorMessage: general.ErrorMessage(e.Field(), e.Tag()),
// 			})
// 			response := model.Response{
// 				ErrorCode: 101,
// 				Message:   "Gagal login, check field yang wajib terisi",
// 				Errors:    errField,
// 			}
// 			return c.JSON(http.StatusOK, response)
// 		}
// 	}
// }
