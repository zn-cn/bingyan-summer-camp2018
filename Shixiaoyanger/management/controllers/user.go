package controllers

import (
//	"fmt"
	"github.com/astaxie/beego"
	"management/models"
	
   // "github.com/astaxie/beego/orm"
	
)

type MainController struct {
	beego.Controller

}

type UserController  struct {
	beego.Controller

}
type ReviewController  struct {
	beego.Controller

}
type RoleController  struct {
	beego.Controller

}

func (this *UserController) First() { 

	this.TplName = "login.html"

	
}

func (this *UserController) Login() { 

	this.TplName = "login.html"

	
}


func (this *ReviewController) Review(){

	u := this.Input().Get("username")
	p := this.Input().Get("password")
	e := this.Input().Get("email")
	n := this.Input().Get("nick")
	err := models.AddUser(u,p,e,n)
	if err != nil {
		beego.Error(err)
	}
  
	
	name :=this.GetString("username")
	pass :=this.GetString("password")
 
	this.Data["password"] = pass
	this.Data["username"] = name
	
	this.Data["Hello"] = this.Ctx.Input.Param(":id")
	
	
	 
	
	this.TplName = "review.html"
}

func (this *UserController) Info(){

	u := this.Input().Get("username")
	p := this.Input().Get("password")
	err1,err2,user := models.FindUserByUserName(u,p)
	if err1&&err2 {
		this.Data["feedback1"] = user.Name
		this.Data["feedback2"] = "登陆成功"
	}else{
		this.Data["feedback1"] = "账号或密码错误，"
		this.Data["feedback2"] = "登陆失败"
	}
	this.TplName = "info.html"
}

func (this *UserController) Register(){

	this.TplName = "regist.html"
}



func (this *UserController) Logout(){


	this.TplName = "regist.html"
}
 






