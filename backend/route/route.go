package route

import (
	"halo_food/config"
	"halo_food/helpers/security"
	levelController "halo_food/modules/level/controllers"
	userController "halo_food/modules/users/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	db := config.ConnectDB()
	e.Use(middleware.Logger())
	e.POST("register/customer", func(c echo.Context) error { return userController.Register(db, c) })
	e.POST("register/driver", func(c echo.Context) error { return userController.RegisterDriver(db, c) })
	e.POST("register/resto", func(c echo.Context) error { return userController.RegisterResto(db, c) })
	e.POST("auth/login", func(c echo.Context) error { return userController.DoLogin(db, c) })

	conf := middleware.JWTConfig{
		Claims:     &security.JwtToken{},
		SigningKey: []byte(config.GetEnv("TOKEN_SECRET")),
	}

	var IsLoggedIn = middleware.JWTWithConfig(conf)
	e.GET("level/get_one_by_id/:id_level", func(c echo.Context) error { return levelController.GetLevelByID(db, c) }, IsLoggedIn)
	e.GET("level/get_all/:limit/:page", func(c echo.Context) error { return levelController.GetAll(db, c) })
}
