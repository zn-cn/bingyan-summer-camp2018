package controllers

import (
	"github.com/astaxie/beego/session"
	"strings"
	"hello/models"
	"fmt"
)

var globalSessions *session.Manager

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

func CheckLogin(c *UserController){
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
