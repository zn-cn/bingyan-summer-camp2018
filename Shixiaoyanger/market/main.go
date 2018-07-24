package main

import (
	_ "market/routers"
	"market/models"

	

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	// 注册数据库
	models.RegisterDB() 
}
func main() {

	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
//	beego.SetStaticPath("/static","static")
	orm.RunSyncdb("default", false, true)  
	
	beego.Run()

}

 