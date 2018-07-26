package main

import (
	"github.com/labstack/echo"
	"project/controller"
)

func main() {
	e := echo.New()

	e.POST("/login", controller.Login)
	e.POST("/sign-up", controller.SignUp)
	e.DELETE("/homepage/members-delete", controller.DeleteMember)
	e.POST("/homepage/groups-show", controller.ShowGroup)
	e.POST("/homepage/information-get", controller.GetInformation)
	e.PUT("/homepage/groups-add", controller.AddGroup)
	e.PUT("/homepage/information-change", controller.ChangeInformation)

	e.Logger.Fatal(e.Start(":8080"))
}
