package main

import (
	"github.com/labstack/echo"
	"project/controller"
)

func main() {
	e := echo.New()
	e.GET("/login", controller.Login)
	e.GET("/signup", controller.SignUp)
	e.GET("/homepage/deletemember", controller.DeleteMember)
	e.GET("/homepage/showgroup", controller.ShowGroup)
	e.GET("/homepage/getinformation", controller.GetInformation)
	e.GET("/homepage/addGroup", controller.AddGroup)
	e.GET("/homepage/changeinformation", controller.ChangeInformation)

	e.Logger.Fatal(e.Start(":8080"))
}