package routers

import (
	"market/controllers"

	"github.com/astaxie/beego"
	
)

//注册路由
func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
	beego.NSNamespace("/users",
		beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
		beego.NSRouter("/update", &controllers.UserController{}, "post:Update"),//update info
	//	beego.NSRouter("/info",&controllers.UserController{},"get:info"),



	//	beego.NSRouter("/passwd", &controllers.UserController{}, "post:Passwd"),
	//头像
	//	beego.NSRouter("/uploads", &controllers.UserController{}, "post:Uploads"),
	//	beego.NSRouter("/downloads", &controllers.UserController{}, "get:Downloads"),
	),
	beego.NSNamespace("/goods",
//		beego.NSRouter("/:id", &controllers.GoodsController{}, "get:GetOne;put:Put;delete:Delete"),
//		beego.NSRouter("/", &controllers.GoodsController{}, "get:GetAll;post:Post"),
//		beego.NSRouter("/auth", &controllers.GoodsController{}, "post:Auth"),
		beego.NSRouter("/test", &controllers.GoodsController{}, "get:Test"),
		beego.NSRouter("/view/:id([0-9]+)", &controllers.GoodsController{}, "get:View"),
		beego.NSRouter("/view/popular", &controllers.GoodsController{}, "get:Popular"),
		beego.NSRouter("/view/latest", &controllers.GoodsController{}, "get:Latest"),
		beego.NSRouter("/add", &controllers.GoodsController{}, "post:Add"),
		beego.NSRouter("/search", &controllers.GoodsController{}, "post:Search"),
	), 
)
beego.AddNamespace(ns)
}
