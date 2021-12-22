package controllers

import (
	"encoding/json"
	"fmt"
	general "halo_food/helpers/general"
	model "halo_food/models"
	drivermodel "halo_food/modules/driver/models"
	restomodel "halo_food/modules/resto/models"
	usermodel "halo_food/modules/users/models"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	userValidatorRegister struct {
		Email        string `json:"email" validate:"required,email"`
		Alias        string `json:"alias" validate:"required"`
		UserPassword string `json:"user_password" validate:"required,min=8"`
		Nama         string `json:"nama" validate:"required,min=3"`
	}

	driverRegister struct {
		Email  string          `json:"email" form:"email" validate:"required,email"`
		Alias  string          `json:"alias" form:"alias" validate:"required"`
		Gender string          `json:"gender" form:"gender" validate:"required"`
		Nama   string          `json:"nama" form:"nama" validate:"required,min=3"`
		NoId   string          `json:"no_id" form:"no_id" validate:"required"`
		Ktp    *multipart.File `json:"ktp" form:"ktp"`
	}

	restoRegister struct {
		Email     string          `json:"email" form:"email" validate:"required,email"`
		Alias     string          `json:"alias" form:"alias" validate:"required"`
		Gender    string          `json:"gender" form:"gender" validate:"required"`
		Nama      string          `json:"nama" form:"nama" validate:"required,min=3"`
		NoId      string          `json:"no_id" form:"no_id" validate:"required"`
		NamaUsaha string          `json:"nama_usaha" form:"nama_usaha" validate:"required,min=5"`
		Alamat    string          `json:"alamat" form:"alamat" validate:"required,min=10"`
		Lokasi    string          `json:"lokasi" form:"lokasi" validate:"required"`
		Ktp       *multipart.File `json:"ktp" form:"ktp"`
	}
)

func Register(db *gorm.DB, c echo.Context) error {
	validate := validator.New()
	user := new(usermodel.Users)
	general.BindData(c, user)

	var response model.Response
	registerValidator := new(userValidatorRegister)
	if err := c.Bind(registerValidator); err != nil {
		return err
	}

	err := validate.Struct(registerValidator)
	if err != nil {
		var errField []model.ErrorValidator
		errs := general.TranslateError(err, validate)
		for idx, e := range err.(validator.ValidationErrors) {
			errField = append(errField, model.ErrorValidator{
				FieldName:    e.Field(),
				ErrorMessage: general.ErrorMessage(e.Field(), errs[idx].Error()),
			})
			response = model.Response{
				ErrorCode: 101,
				Message:   "Gagal melakukan registrasi, check field yang wajib terisi",
				Errors:    errField,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	user.ConfirmKey = general.RandomString(20)
	newPass, _ := bcrypt.GenerateFromPassword([]byte(registerValidator.UserPassword), bcrypt.DefaultCost)
	user.UserPassword = string(newPass)
	user.IDLevel = 3

	err = user.RegisterNewCustomer(db)
	if err != nil {
		response = model.Response{
			ErrorCode: 101,
			Message:   "Gagal melakukan registrasi pengguna ",
			Errors:    err,
		}
	} else {
		response = model.Response{
			ErrorCode: 0,
			Message:   "Sukses melakukan registrasi pengguna",
			Data:      user,
		}
	}

	return c.JSON(http.StatusOK, response)

}

func RegisterDriver(db *gorm.DB, c echo.Context) error {
	validate := validator.New()
	user := new(usermodel.Users)
	general.BindData(c, user)

	var response model.Response

	driverRegister := new(driverRegister)
	if err := c.Bind(driverRegister); err != nil {
		return err
	}

	err := validate.Struct(driverRegister)
	if err != nil {
		var errField []model.ErrorValidator
		errs := general.TranslateError(err, validate)
		for idx, e := range err.(validator.ValidationErrors) {
			errField = append(errField, model.ErrorValidator{
				FieldName:    e.Field(),
				ErrorMessage: general.ErrorMessage(e.Field(), errs[idx].Error()),
			})
		}
		response = model.Response{
			ErrorCode: 101,
			Message:   "Gagal melakukan registrasi, periksa field yang wajib terisi",
			Errors:    errField,
		}
		return c.JSON(http.StatusOK, response)
	}
	user.ConfirmKey = general.RandomString(20)
	user.IDLevel = 5
	tx := db.Begin()
	err = user.RegisterNewCustomer(tx)
	if err != nil {
		response = model.Response{
			ErrorCode: 102,
			Message:   "Gagal melakukan registrasi pengguna ",
			Errors:    err,
		}
		tx.Rollback()
	} else {
		ktpFile, header, fileError := c.Request().FormFile("ktp")
		if fileError != nil {
			tx.Rollback()
			response = model.Response{
				ErrorCode: 102,
				Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
				Errors:    err,
			}
			return c.JSON(http.StatusOK, response)
		}
		if filepath.Ext(header.Filename) == ".jpg" {
			out, pathError := ioutil.TempFile("files", "ktp-*.jpg")
			if pathError != nil {
				tx.Rollback()
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
					Errors:    err,
				}
				return c.JSON(http.StatusOK, response)
			}
			defer out.Close()
			_, fileName := filepath.Split(out.Name())
			_, copyError := io.Copy(out, ktpFile)
			if copyError != nil {
				log.Println("Error copying", copyError)
			}

			if err != nil {
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
					Errors:    err,
				}
				tx.Rollback()
			}
			driver := &drivermodel.Driver{
				IDUser: user.IdUser,
				NoId:   c.FormValue("no_id"),
				NoPlat: c.FormValue("no_plat"),
				Ktp:    fileName,
			}
			err = driver.RegisterNewDriver(tx)
			if err != nil {
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi pengguna ",
					Errors:    err,
				}
				tx.Rollback()
			} else {
				response = model.Response{
					ErrorCode: 0,
					Message:   "Sukses melakukan registrasi",
				}
				tx.Commit()
			}
		} else {
			response = model.Response{
				ErrorCode: 102,
				Message:   "Gagal melakukan registrasi pengguna, file gambar tidak didukung",
			}
			tx.Rollback()
		}
	}
	return c.JSON(http.StatusOK, response)
}

func RegisterResto(db *gorm.DB, c echo.Context) error {
	validate := validator.New()
	user := new(usermodel.Users)
	general.BindData(c, user)

	var response model.Response

	restoRegister := new(restoRegister)
	if err := c.Bind(restoRegister); err != nil {
		return err
	}

	fmt.Println(restoRegister)
	err := validate.Struct(restoRegister)
	if err != nil {
		var errField []model.ErrorValidator
		errs := general.TranslateError(err, validate)
		for idx, e := range err.(validator.ValidationErrors) {
			errField = append(errField, model.ErrorValidator{
				FieldName:    e.Field(),
				ErrorMessage: general.ErrorMessage(e.Field(), errs[idx].Error()),
			})
		}
		response = model.Response{
			ErrorCode: 101,
			Message:   "Gagal melakukan registrasi, periksa field yang wajib terisi",
			Errors:    errField,
		}
		return c.JSON(http.StatusOK, response)
	}
	user.ConfirmKey = general.RandomString(20)
	user.IDLevel = 5
	tx := db.Begin()
	err = user.RegisterNewCustomer(tx)
	if err != nil {
		response = model.Response{
			ErrorCode: 102,
			Message:   "Gagal melakukan registrasi pengguna ",
			Errors:    err,
		}
		tx.Rollback()
	} else {
		ktpFile, header, fileError := c.Request().FormFile("ktp")
		if fileError != nil {
			tx.Rollback()
			response = model.Response{
				ErrorCode: 102,
				Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
				Errors:    err,
			}
			return c.JSON(http.StatusOK, response)
		}
		if filepath.Ext(header.Filename) == ".jpg" {
			out, pathError := ioutil.TempFile("files", "ktp-*.jpg")
			if pathError != nil {
				tx.Rollback()
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
					Errors:    err,
				}
				return c.JSON(http.StatusOK, response)
			}
			defer out.Close()
			_, fileName := filepath.Split(out.Name())
			_, copyError := io.Copy(out, ktpFile)
			if copyError != nil {
				log.Println("Error copying", copyError)
			}

			if err != nil {
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi harap mengunggah scan ktp",
					Errors:    err,
				}
				tx.Rollback()
			}
			var location map[string]interface{}
			err = json.Unmarshal([]byte(restoRegister.Lokasi), &location)
			fmt.Println(location)
			if err != nil {
				return err
			}
			resto := &restomodel.Resto{
				IdUser:    user.IdUser,
				NoId:      restoRegister.NoId,
				NamaUsaha: restoRegister.NamaUsaha,
				Alamat:    restoRegister.Alamat,
				Lokasi:    model.Location{X: location["X"].(float64), Y: location["Y"].(float64)},
				Ktp:       fileName,
			}
			err = resto.RegisterNewResto(tx)
			fmt.Println(err)
			if err != nil {
				response = model.Response{
					ErrorCode: 102,
					Message:   "Gagal melakukan registrasi pengguna ",
					Errors:    err,
				}
				tx.Rollback()
			} else {
				response = model.Response{
					ErrorCode: 0,
					Message:   "Sukses melakukan registrasi",
				}
				tx.Commit()
			}
		} else {
			response = model.Response{
				ErrorCode: 102,
				Message:   "Gagal melakukan registrasi pengguna, file gambar tidak didukung",
			}
			tx.Rollback()
		}
	}
	return c.JSON(http.StatusOK, response)
}
