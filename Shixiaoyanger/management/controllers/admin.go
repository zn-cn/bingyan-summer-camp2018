package controllers

import (
//	"fmt"
	"github.com/astaxie/beego"
	"management/models"
	
   // "github.com/astaxie/beego/orm"
	
)


type AdminController  struct {
	beego.Controller

}


func (this *AdminController) Admin(){
	
	this.TplName = "admin.html"
}

func (this *AdminController) Manage(){
    this.TplName = "manage.html"
// op := this.Input().Get("op")
	user,err :=models.GetAllUser()
	this.Data["user"] =user

	if err != nil {
		beego.Error(err)
	}
	

	id := this.Input().Get("id")
	if len (id) ==0 {
		return
	}
  err =models.DeleteUser(id)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/manage",302)

}


func (this *AdminController) Status(){
this.TplName ="status.html"
u := this.Input().Get("username")
p := this.Input().Get("password")
err1,err2,user := models.FindUserByUserName(u,p)
if err1&&err2 {
	this.Data["feedback1"] = user.Name
	this.Data["feedback2"] = "登陆成功"
	this.Redirect("/admin/manage",302)
}else{
	this.Data["feedback1"] = "账号或密码错误，"
	this.Data["feedback2"] = "登陆失败"
}
}