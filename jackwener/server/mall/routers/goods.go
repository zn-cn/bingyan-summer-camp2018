package routers

import (
	"mall/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/api/goods/kind", &controllers.MainController{}, "post:KindSearch")
	beego.Router("/api/goods/local", &controllers.MainController{}, "post:LocalSearch")
	beego.Router("/api/goods/:good", &controllers.MainController{}, "get:Goods")
	beego.Router("/api/goods/:good/picture", &controllers.MainController{}, "post:Picture")
	beego.Router("/api/goods/recent", &controllers.MainController{}, "get:RecentSearch")
	beego.Router("/api/goods/popular", &controllers.MainController{}, "get:PopularSearch")
}