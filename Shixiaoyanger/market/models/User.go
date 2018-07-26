package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
func RegisterDB(){
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:123456@/market?charset=utf8")

    orm.RegisterModel(new(User),new(Goods) )
}
 

type User struct{
	Id        int64  // `json:"-"` 
    Username  string   `json:"Username"`
    Password  string   `json:"password"`
    Tel       string   `json:"tel"`
	Views     int64    `json:"views"`
	
}
type Users struct{
	User *User


}
//locate User by username
func ReadUser(user *User) error{
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err := qs.Filter("username",user.Username).One(user)
	return err
	//err == nil username has existed
	//err!= nil  .....not exist

}

//注册添加数据成员
func AddUser(user *User)  (bool,error){
	o := orm.NewOrm()
	err := ReadUser(user)
	if err == nil{
		return false, err //username has existed
	}
	user.Password = GetDes(user.Password)
	_,err =o.Insert(user)
	if err != nil{
		return false, err  //faild to insert
	}
	return true, nil
}

func UpdateUser(user *User) error{
	o := orm.NewOrm()
	err := ReadUser(user)
	if err != nil{
		return err//not exist
	}
	//num,err1 :=o.Update(&user,"Tel",)
	num, err1 := o.QueryTable("user").Update(orm.Params{
		"views": orm.ColValue(orm.ColAdd, 100),
	})
	fmt.Println("num",num)
	return err1 //update success
}

func IncreaseView(user *User) {
	o := orm.NewOrm()
	o.QueryTable("user").Update(orm.Params{
		"views": orm.ColValue(orm.ColAdd, 1),
	})
}

// search User, return info
func FindbyUserame(username string) (bool, User){
	o :=orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("username", username).One(&user)
	return err!= orm.ErrNoRows, user
}

func DeleteUser(user *User) error {
	o := orm.NewOrm()
	_,err := o.Delete(user)
	if err!= nil{
		return err
	}
	return nil
}