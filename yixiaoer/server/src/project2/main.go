package main

import (
	"github.com/labstack/echo"
	"project/project2/controller"
)

func main() {
	e := echo.New()

	e.POST("/login", controller.Login)
	e.POST("/sign-up", controller.SignUp)
	e.POST("/page/categories", controller.ShowCategory)
	e.POST("/page/location", controller.ShowLocation)
	e.POST("/page/commodities", controller.CommodityInfo)
	e.POST("/page/popularity", controller.PopluarRank)
	e.POST("/homepage",controller.UserInfo)

	e.Logger.Fatal(e.Start(":8080"))
}
