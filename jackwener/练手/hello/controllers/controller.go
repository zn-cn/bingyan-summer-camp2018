package controllers

import (
	"hello/models/class"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation" // 用于校验信息
)

type UserController struct {
	beego.Controller
}

func (c *UserController) PageLogin() {
	c.TplName = "login.html" // 将hello.html页面输出
}
func (c *UserController) Login() {
	name := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("This is name and password")
	valid := validation.Validation{}
	valid.Required(name, "username") // 校验是否为空值
	valid.Required(password, "password")
	// valid.MaxSize(id, 20, "id")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "error.html"
		return
	}
	fmt.Println(name, password)
	u := class.User{
		Name: name,
		Password:password,
	}
	err := u.ReadDB()
	if err != nil {
		fmt.Println(err)
		c.TplName = "error.html"
		return
	} else {
		c.TplName = "success.html"
	}

}

func (c *UserController) PageRegister() {
	c.TplName = "register.html"
}

func (c *UserController) Register() {
	name := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("This is name and password")
	fmt.Println(name, password)

	valid := validation.Validation{}
	valid.Required(name, "username") // 校验是否为空值
	valid.Required(password, "password")
	// valid.MaxSize(id, 20, "id")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "error.html"
		return
	}

	u := class.User{
		Name:      name,
		Password: password,
	}

	err := u.Create()
	if err != nil {
		fmt.Println(err)
		c.TplName = "error.html"
		return
	}else {
		c.TplName = "success.html"
	}
}

