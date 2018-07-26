package routers

import (
	"mall/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/person", &controllers.MainController{}, "post:Login")
	beego.Router("/api/person/register", &controllers.MainController{}, "post:Register")
	beego.Router("/api/person/:user", &controllers.MainController{}, "get:Person;post:Login")
}