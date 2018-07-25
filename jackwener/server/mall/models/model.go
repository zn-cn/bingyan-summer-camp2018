package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Good struct {
	Id      string    `orm:"pk"`
	Title   string
	Local   string
	Kind    string
	Intro   string
	Price   int
	Cars  	[]*Car   `orm:"rel(m2m)"`
	Views   int
}

type Car struct {
	Id       string   `orm:"pk"`
	User     *User  `orm:"reverse(one)"`
	Goods    []*Good `orm:"reverse(many)"`
}

type User struct{
	Id       string  `orm:"pk"`
	Name     string
	Password string
	Nickname string
	PageViews int
	Car      *Car `orm:"rel(one)"`
	Inform   []*Inform  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Record struct {
	Id      int  `orm:"pk"`
	Content string
	Date    orm.DateTimeField
	Number  int
}

type  Inform struct {
	Id       string   `orm:"pk"`
	User     *User  `orm:"rel(fk)"`
	Content  string
}

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:jakevin@tcp(localhost:3306)/mall_goods?charset=utf8")
	orm.RegisterModel(new(User), new(Good),new(Car),new(Record),new(Inform)) // 注册模型，建立User类型对象，注册模型时，需要引入包
	orm.RunSyncdb("default",false,true)
	orm.DefaultTimeLoc = time.UTC
}