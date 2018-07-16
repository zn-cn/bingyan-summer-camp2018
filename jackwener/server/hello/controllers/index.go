package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"hello/encryptions"
	"hello/models"
	"hello/models/class"
	"github.com/astaxie/beego/orm"
)


type UserController struct {
	beego.Controller
}

func (c *UserController) PageLogin() {

}

//通过路由"/"，发送post请求，默认Json数据中User[0]的Name,Password为发送的信息值，返回的Json数据中Result:false与true说明了登录情况
func (c *UserController) Login() {
	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	name := ob.User[0].Name
	pwd := ob.User[0].Password
	password := encryptions.Salt(pwd)//MD5加盐加密
	u := class.User{
		Name: name,
		Password:password,
	}
	fmt.Print(u)
	o := orm.NewOrm()
	err := o.QueryTable("user").Exclude("status", "apply").Filter("name", name).Filter("password", password).One(&u)
	permission := u.Status
	//var str string = strconv.Itoa(id)
	if err == nil {
		fmt.Println(err)
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
		c.SetSession("userPermission", permission)
		fmt.Print("设置了")
		return
	} else {
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
	}
}

//通过路由"/"，发送post请求，默认Json数据中User[0]的Name,Password为发送的信息值，注册成功/失败Json中result相应为true/false
func (c *UserController) Register() {
	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)

	u := class.User{
		Name:      ob.User[0].Name,
		Password: encryptions.Salt(ob.User[0].Password),
		Nickname: ob.User[0].Nickname,
		Email:    ob.User[0].Email,
		Phone :   ob.User[0].Phone,
		Group :   ob.User[0].Group,
		Status :  "apply",
	}
	fmt.Print(u)
	err := u.Create()
	if err == nil {
		fmt.Println(err)
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	}
}