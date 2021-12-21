package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
)

func ErrorMessage(field string, tag string) string {
	switch tag {
	case "required":
		return field + " tidak boleh kosong!"
	case "email":
		return field + " harus berupa email!"
	default:
		return tag
	}
}

func BindData(c echo.Context, data interface{}) {
	contentTypes := c.Request().Header.Get("Content-Type")
	if contentTypes == "application/json" {
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		json.Unmarshal(bodyBytes, data)

	} else {
		c.Bind(data)
	}
}

func TranslateError(err error, validate *validator.Validate) (errs []error) {

	indo := id.New()
	uni := ut.New(indo, indo)
	transl, _ := uni.GetTranslator("id")
	_ = id_translations.RegisterDefaultTranslations(validate, transl)

	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(transl))
		errs = append(errs, translatedErr)
	}
	return errs
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
