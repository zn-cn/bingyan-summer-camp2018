package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func IndexInit() {
	beego.Router("/", &controllers.UserController{}, "get:PageLogin;post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
}