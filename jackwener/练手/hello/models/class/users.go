package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id       int `orm:"pk"` // 设置为主键，字段Id, Password首字母必须大写
	Name     string
	Password string
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	fmt.Println(*u)
	err = o.Read(u, "name", "password")
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	fmt.Println("Create success!")
	fmt.Println(*u)
	_, _ = o.Insert(u)
	return err
}

func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}


