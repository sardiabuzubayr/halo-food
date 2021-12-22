package main

import (
	"halo_food/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	route.Init(e)
	e.Start(":9000")
}
