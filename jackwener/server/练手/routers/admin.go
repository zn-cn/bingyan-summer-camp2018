package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func AdminInit() {
	beego.Router("/manager/apply", &controllers.UserController{}, `get:ListApply;post:UpdateApply`)
	beego.Router("/manager/user", &controllers.UserController{}, `get:UserList`)
	beego.Router("/manager/user/add", &controllers.UserController{}, `get:PageUserAdd;post:UserAdd`)
	beego.Router("/manager/user/delete", &controllers.UserController{}, `get:PageUserDelete;post:UserDelete`)
	beego.Router("/manager/user/update", &controllers.UserController{}, `get:PageUserUpdate;post:UserUpdate`)
	beego.Router("/manager/user/group", &controllers.UserController{}, `get:PageGroupList;post:GroupList`)
}