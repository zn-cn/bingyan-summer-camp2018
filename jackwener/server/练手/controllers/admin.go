package controllers

import (
	"hello/models/class"
	"github.com/astaxie/beego/orm"
	"fmt"
	"hello/models"
	"encoding/json"
	"hello/encryptions"
	"strings"
)

// 通过路由/manager/apply，发送get请求就能返回所有身份为"apply"(申请者)的Json数据
func (c *UserController) ListApply(){
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var posts []class.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("status","apply").All(&posts, "id", "name", "nickname", "email", "phone", "group", "status")
	fmt.Println(posts)
	u := models.Info{User:posts,Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

// 通过路由/manager/apply，发送post请求，通过请求中Json的User[]的id,可以批量批准Json中Use[]成员的注册
// 即将"status"由"apply"改为"stuff"
func (c *UserController) UpdateApply() {
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	users := ob.User
	o := orm.NewOrm()
	for _, user := range users{
		id := user.Id
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"status": "stuff"})
	}
	u := models.Info{Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

// 通过路由/manager/user,发送get请求，就能返回所有非"apply"成员的信息的Json数据
func (c *UserController) UserList() {
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var posts []class.User
	o := orm.NewOrm()
	o.QueryTable("user").Exclude("status","apply").All(&posts, "id", "name", "nickname", "email", "phone", "group", "status")
	fmt.Println(posts)
	u := models.Info{User:posts,Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

func(c *UserController) PageUserAdd() {
	CheckLogin(c)
}

// 通过路由/manager/user/add,发送post请求，批量添加Json中User[]的用户
// 管理员添加用户默认为"staff"，添加不可指定id，避免冲突
func(c *UserController) UserAdd() {
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	users := ob.User
	for _, user := range users {
		user.Status = "staff"
		user.Password = encryptions.Salt(user.Password)
	}
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	i, _ := qs.PrepareInsert()
	for _, user := range users {
		id, err := i.Insert(&user)
		if err == nil {
			fmt.Println(id)
		}
	}
	i.Close() // 不要忘记关闭 statement
	u := models.Info{Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) PageUserDelete(){
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

}

// 通过路由/manager/user/delete,发送post请求，通过读取请求Json中User[]的id值，批量删除Json中User[]的用户
func (c *UserController) UserDelete(){
	CheckLogin(c)
	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	o := orm.NewOrm()
	users := ob.User
	for _, user := range users {
		id := user.Id
		if num, err := o.QueryTable("user").Filter("id", id).Delete(); err == nil {
			fmt.Println(num)
		}
	}
	u := models.Info{Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) PageUserUpdate() {
	CheckLogin(c)
}

// 通过路由/manager/user/update,发送post请求，默认通过读取请求Json中User[0]的id值，单个更新Json中User[]的用户
// 不可更新id值
func (c *UserController) UserUpdate() {
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	o := orm.NewOrm()
	user := ob.User[0]
	id := user.Id
	if user.Status !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"status": user.Status})
	}
	if user.Name !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"name": user.Name})
	}
	if user.Nickname !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"nickname": user.Nickname})
	}
	if user.Group !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"group": user.Group})
	}
	if user.Password !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"password": encryptions.Salt(user.Password)})
	}
	if user.Email !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"email": user.Email})
	}
	if user.Phone !="" {
		o.QueryTable("user").Filter("id", id).Update(orm.Params{"phone": user.Phone})
	}
	u := models.Info{Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) PageGroupList(){
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

}

// 通过路由/manager/user/group,发送post请求，默认通过读取请求Json中User[0]的group值，列出该group中所以成员
func (c *UserController) GroupList() {
	err := c.GetSession("userPermission")
	fmt.Print(err)
	if err == nil{
		fmt.Printf("没有session")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else if !strings.Contains(c.GetSession("userPermission").(string), "admin") {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("没权限")
		u := models.Info{}
		u.Result = false
		c.Data["json"] = u
		c.ServeJSON()
		return
	} else {
		fmt.Printf(c.GetSession("userPermission").(string))
		fmt.Printf("有权限")
		u := models.Info{}
		u.Result = true
		c.Data["json"] = u
		c.ServeJSON()
	}

	var ob models.Info
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	group := ob.User[0].Group
	var posts []class.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("group", group).Exclude("status","apply").All(&posts, "id", "name", "nickname", "email", "phone", "group", "status")
	fmt.Println(posts)
	u := models.Info{User:posts,Result:true}
	c.Data["json"] = u
	c.ServeJSON()
}