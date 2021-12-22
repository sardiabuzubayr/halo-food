package controllers

import (
	"halo_food/config"
	general "halo_food/helpers/general"
	"halo_food/helpers/security"
	model "halo_food/models"
	usermodel "halo_food/modules/users/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	UserValidator struct {
		Username string `json:"username" form:"username" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}
)

func DoLogin(db *gorm.DB, c echo.Context) error {
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
				Message:   "Gagal login, mohon cek kembali username dan password anda",
				Errors:    errField,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	if usermodel.DoLoginUser(db, userValidator.Username, userValidator.Password) {
		user, _ := usermodel.GetOneByUsername(db, userValidator.Username)

		lifeTime, _ := strconv.Atoi(config.GetEnv("TOKEN_LIFETIME"))
		lifeTimeRefresh, _ := strconv.Atoi(config.GetEnv("TOKEN_REFRESH_LIFETIME"))
		dataToken := &security.JwtToken{
			Email:     user.Email,
			Alias:     user.Alias,
			Nama:      user.Nama,
			IdLevel:   int(user.Level.IdLevel),
			NamaLevel: user.Level.NamaLevel,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(lifeTime)).Unix(),
			},
		}

		dataRefreshToken := &security.JwtRefreshToken{
			Email: user.Email,
			Alias: user.Alias,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(lifeTimeRefresh)).Unix(),
			},
		}
		token, _ := security.EncodeToken(dataToken)
		refreshToken, _ := security.EncodeToken(dataRefreshToken)

		response := model.Response{
			ErrorCode:    101,
			Message:      "Sukses login",
			Token:        token,
			RefreshToken: refreshToken,
		}
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusOK, false)
}

// func GetAll(c echo.Context) error {
// 	limit, _ := strconv.Atoi(c.Param("limit"))
// 	page, _ := strconv.Atoi(c.Param("page"))
// 	keywords := c.QueryParam("keywords")
// 	data := filmModel.GetAll(db, limit, page, keywords)

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
