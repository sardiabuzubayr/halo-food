package route

import (
	levelController "halo_food/modules/level/controllers"
	userController "halo_food/modules/users/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("auth/login", userController.DoLogin)
	e.GET("level/get_one_by_id/:id_level", levelController.GetLevelByID)
	e.GET("level/get_all/:limit/:page", levelController.GetAll)
	e.POST("register/customer", userController.Register)
	e.POST("register/driver", userController.RegisterDriver)
	e.POST("register/resto", userController.RegisterResto)
}
