package controllers

import (
	"hello/encryptions"
	"github.com/astaxie/beego/validation"
	"fmt"
	"hello/models/class"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

// var globalSessions *session.Manager
/*
func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}
*/

type UserController struct {
	beego.Controller
}

func (c *UserController) PageLogin() {
	/*
	_, err := c.CheckLogin()
	if err {
		c.Redirect("/", 302)
		c.StopRun()
	}
	*/
	c.TplName = "login.html"
	//c.TplName = "index.tpl"

}
/*
func (c *UserController) CheckLogin() (interface{}, bool) {
	userData := c.GetSession("LoginUser")
	if userData == nil {
		return nil, false
	}
	return userData, true
}
*/
/*
func (this *UserController) Session(id int) {
	//根据用户ID查询出角色
	//result := class.RbacRoleUser{UserId: id}
	o := orm.NewOrm()
	//查询出角色信息
	roleResult := class.RbacRoleUser{Id: id}
	o.Read(&roleResult)
	this.SetSession("RoleInfo", roleResult)
	//根据角色查询出权限
	var accessResult []orm.Params
	o.Raw("select t1.id,t1.name,t1.title,t1.sort,t1.pid,t1.level,t.role_id" +
		" from rbac_access t inner join rbac_node t1" +
		" on t.node_id=t1.id " +
		" where t1.status=1 and t.role_id='" + strconv.Itoa(roleResult.Id) + "' and t1.is_show=1 ").Values(&accessResult)

	leftTreeResult = []orm.Params{}
	tmpResult := this.TreeNodeRecursion(accessResult, 0)
	this.SetSession("LeftNavResult", tmpResult)
}
*/
func (c *UserController) Login() {
	name := c.GetString("username")
	pwd := c.GetString("password")
	password := encryptions.Salt(pwd)
	valid := validation.Validation{}
	valid.Required(name, "username")
	valid.Required(password, "pwd")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "error.html"
		return
	}
	u := class.User{
		Name: name,
		Password:password,
	}
	result := class.User{}
	err := u.ReadDB()
	if err != nil {
		fmt.Println(err)
		c.TplName = "error.html"
		return
	} else {
		c.SetSession("LoginUser", result)
		//c.SessionRbac(result.Id)
		c.TplName = "manager.html"
	}
}

func (c *UserController) PageRegister() {
	c.TplName = "register.html"
}

func (c *UserController) Register() {
	name := c.GetString("username")
	pwd := c.GetString("password")
	password := encryptions.Salt(pwd)
	phone := c.GetString("phone")
	email := c.GetString("email")
	nickname := c.GetString("nickname")
	group := c.GetString("group")
	fmt.Println("This is name and password")
	fmt.Println(name, password)

	valid := validation.Validation{}
	valid.Required(name, "username")
	valid.Required(pwd, "password")
	valid.Required(nickname, "nickname")
	valid.Required(email, "email")
	valid.Required(phone, "phone")
	valid.Required(group, "group")

	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "error.html"
		return
	}

	u := class.User{
		Name:      name,
		Password: password,
		Nickname: nickname,
		Email:    email,
		Phone :   phone,
		Group :   group,
		Status :  "apply",
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

func (c *UserController) UserList() {
	c.TplName = "superadmin/user.html"
	var posts []class.User
	o := orm.NewOrm()
	o.QueryTable("user").Exclude("status","apply").All(&posts, "id", "name", "nickname", "email", "phone", "group", "status")
	fmt.Println(posts)
	c.Data["posts"] = posts
}

func (c *UserController) UserUpdate() {
	c.TplName = "superadmin/user.html"
	u := class.User{}
	id, _ :=c.GetInt("id")
	// ...
	if c.GetString("username") != "" {
		name := c.GetString("username")
		u.Name = name
	}
	if c.GetString("phone") != "" {
		phone := c.GetString("phone")
		u.Phone = phone
	}
	if c.GetString("email") != "" {
		email := c.GetString("email")
		u.Email = email
	}
	if c.GetString("nickname") != "" {
		nickname := c.GetString("nickname")
		u.Nickname = nickname
	}
	if c.GetString("group") != "" {
		group := c.GetString("group")
		u.Group = group
	}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("id",id).Update(orm.Params{"name":u.Name,"phone":u.Phone,"email":u.Email,"nickname":u.Nickname,"group":u.Group})
}

func (c *UserController) PageUserAdd() {
	c.TplName = "superadmin/add.html"
}

func (c *UserController) UserAdd() {
	name := c.GetString("username")
	pwd := c.GetString("password")
	password := encryptions.Salt(pwd)
	phone := c.GetString("phone")
	email := c.GetString("email")
	nickname := c.GetString("nickname")
	group := c.GetString("group")
	valid := validation.Validation{}
	valid.Required(name, "username")
	valid.Required(pwd, "password")
	valid.Required(nickname, "nickname")
	valid.Required(email, "email")
	valid.Required(phone, "phone")
	valid.Required(group, "group")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "error.html"
		return
	}
	u := class.User{
		Name:      name,
		Password: password,
		Nickname: nickname,
		Email:    email,
		Phone :   phone,
		Group :   group,
		Status :  "admin",
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

func (c *UserController) UserDelete(){
	id, _ :=c.GetInt("id")
	u := class.User{Id:id}
	o := orm.NewOrm()
	if num, err := o.Delete(u); err == nil {
		fmt.Println(num)
	}
}

func (c *UserController) Approve(){
	c.TplName = "superadmin/approve.html"
	id, _ :=c.GetInt("id")
	o := orm.NewOrm()
	o.QueryTable("user").Filter("id",id).Update(orm.Params{"status":"admin"})

}

func (c *UserController) PageApprove(){
	c.TplName = "superadmin/approve.html"
	var posts []class.User
	o := orm.NewOrm()
	o.QueryTable("user").Exclude("status","apply").All(&posts, "id", "name", "nickname", "email", "phone", "group", "status")
	fmt.Println(posts)
	c.Data["posts"] = posts
}
