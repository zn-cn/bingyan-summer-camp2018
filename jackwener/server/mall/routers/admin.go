package routers

import (
	"mall/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/api/admin/kind", &controllers.MainController{}, "post:KindSearch")
	beego.Router("/api/admin/local", &controllers.MainController{}, "post:LocalSearch")
}