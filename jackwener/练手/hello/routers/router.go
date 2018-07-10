package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{},`get:PageLogin`)
	beego.Router("/", &controllers.UserController{}, `post:Login`)
	beego.Router("/login", &controllers.UserController{}, `get:PageLogin`)
	beego.Router("/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/register", &controllers.UserController{}, `get:PageRegister`)
	beego.Router("/register", &controllers.UserController{}, `post:Register`)

}