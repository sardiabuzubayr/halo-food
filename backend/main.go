package main

import (
	"halo_food/config"
	"halo_food/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	e := echo.New()
	route.Init(e)
	e.Start(":9000")
}
