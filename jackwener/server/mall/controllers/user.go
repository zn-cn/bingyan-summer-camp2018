package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"fmt"
	"mall/encryptions"
	"mall/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Login() {
	var ob models.UserJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	name := ob.User.Name
	pwd := ob.User.Password
	password := encryptions.Salt(pwd)//MD5加盐加密
	user := models.User{}
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).Filter("password", password).One(&user)
	id := user.Id
	fmt.Print(err)
	if err == nil {
		fmt.Print(err)
		var reJson models.UserJson
		reJson.Status = 200
		reJson.Message = "登录成功"
		c.Data["json"] = reJson
		c.ServeJSON()
		c.SetSession("userId", id)
		fmt.Print("设置了")
		return
	} else {
		var reJson models.UserJson
		reJson.Status = 401
		reJson.Message = "登录失败"
		c.Data["json"] = reJson
		c.ServeJSON()
	}
}

// 注册
func (c *MainController) Register() {
	var ob models.UserJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	name := ob.User.Name
	pwd := ob.User.Password
	password := encryptions.Salt(pwd)
	nickname := ob.User.Nickname
	o := orm.NewOrm()
	user := models.User{
		Id:        name,
		Name:      name,
		Password: password,
		Nickname: nickname,
		PageViews: 0,
	}
	var car models.Car
	car.Id = name
	user.Car = &car
	o.Insert(&car)
	_,err := o.Insert(&user)
	if err == nil {
		reJson := models.UserJson{}
		reJson.Status = 200
		reJson.Message = "注册成功"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	} else {
		fmt.Print(err)
		reJson := models.UserJson{}
		reJson.Status = 406
		reJson.Message = "注册失败"
		c.Data["json"] = reJson
		c.ServeJSON()
	}
}

// 返回个人信息页的一些基本信息
func (c *MainController) Person() {
	name := c.GetString(":user")
	fmt.Printf(name)
	var user models.User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).One(&user)
	if err != nil {
		var reJson models.UserJson
		reJson.Status = 400
		reJson.Message = "用户不存在"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
	o.QueryTable("user").Filter("name", name).Update(orm.Params{"page_views": orm.ColValue(orm.ColAdd, 1)})
	var reJson models.UserJson
	reJson.User =user
	reJson.Status = 200
	reJson.Message = "返回成功"
	c.Data["json"] = reJson
	c.ServeJSON()
	return
}

