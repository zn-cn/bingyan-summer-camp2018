package routers

import (
	"management/controllers"
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)



func init(){


	beego.Router("/", &controllers.UserController{} ,"get:First")
	beego.Router("/v1/register", &controllers.UserController{}, "get:Register")

   beego.Router("/review", &controllers.ReviewController{},"get:Review") 

   beego.Router("/user", &controllers.UserController{},"get:Info") 
   beego.Router("/admin", &controllers.AdminController{},"get:Admin") 
   beego.Router("/admin/status", &controllers.AdminController{},"get:Status")
   beego.Router("/admin/manage", &controllers.AdminController{},"get:Manage") 


   ns := beego.NewNamespace("/v1",
     beego.NSNamespace("/users",
	 //  beego.NSRouter("/register", &controllers.UserController{}, "get:Register"),
	   beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
	   beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
//	   beego.NSRouter("/passwd", &controllers.UserController{}, "post:Passwd"),
//	   beego.NSRouter("/uploads", &controllers.UserController{}, "post:Uploads"),
//	   beego.NSRouter("/downloads", &controllers.UserController{}, "get:Downloads"),
	),
	/*
      beego.NSNamespace("/admin",
	      beego.NSRouter("/:id", &controllers.AdminController{}, "get:GetOne;put:Put;delete:Delete"),
	      beego.NSRouter("/", &controllers.AdminController{}, "get:GetAll;post:Post"),
	      beego.NSRouter("/auth", &controllers.AdminController{}, "post:Auth"),
   ),
   */
)
beego.AddNamespace(ns)



}


