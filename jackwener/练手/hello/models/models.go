package models

import(
	"hello/models/class" // 注册模型，需要引入该包
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
使用orm连接数据库步骤：
//告诉orm使用哪一种数据库
1.注册数据库驱动RegisterDriver(driverName, DriverType)
2.注册数据库RegisterDataBase(aliasName, driverName, dataSource, params ...)
3.注册对象模型RegisterModel(models ...)
4.开启同步RunSyncdb(name string, force bool, verbose bool)
*/

// 在init函数中连接数据库，当导入该包的时候便执行此函数
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:jakevin@tcp(localhost:3306)/project?charset=utf8")
	orm.RegisterModel(new(class.User)) // 注册模型，建立User类型对象，注册模型时，需要引入包
	orm.RunSyncdb("default", false, true)
}