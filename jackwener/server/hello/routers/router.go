package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "get:PageLogin;post:Login")
	beego.Router("/login", &controllers.UserController{}, "get:PageLogin;post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register;get:PageRegister")
	beego.Router("/user", &controllers.UserController{}, `get:UserList;post:UserUpdate`)
	beego.Router("/add", &controllers.UserController{}, `get:PageUserAdd;post:UserAdd`)
	//beego.Router("/delete", &controllers.UserController{}, `get:UserDelete`)
	beego.Router("/approve", &controllers.UserController{}, `get:PageApprove;post:Approve`)
}