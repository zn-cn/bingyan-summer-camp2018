package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	//"project/controller"不可以循环导包
)

type User struct {
	Id       string `json:"id" form:"id"query:"id"`
	Password string `json:"password" form:"password"query:"password"`
	Email    string `json:"email" form:"email" query:"email"`
	Phone    string `json:"phone" form:"phone" query:"phone"`
	Name     string `json:"name" form:"name" query:"name"`
	Group    string `json:"group" form:"group" query:"group"`
	Identity string `json:"identity" form:"identity" query:"identity"`
	Status   string `json:"status" form:"status" query:"status"`
}

func Login(u map[string]string) int8 {
	var i int8
	if u ["name"]!= "" && u["password"]!="" {

		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true) //连接数据库

		c := session.DB("test").C("people")

		var user1 User
		var user2 User
		c.Find(bson.M{"name": u["name"]}).One(&user1) //id由数据库分配，不方便登陆，用name登陆
		if user1.Name != "" { //查找是否存在这个name的用户
			c.Find(bson.M{"name": u["name"],
				"password": u["password"]}).One(&user2)
			if user2.Name != "" {
				i = 0 //若存在查找是否这个name的用户pw也一致
			} else if user2.Name == "" {
				i = 1 //密码错误
			}
		} else if user1.Name == "" {
			i = 2 //没有这个name的用户
		}




	} else {
		i = 2
	}
	return i
}

func SignUp(u map[string]string) int8 {
	var i int8
	println(u)
	if u["password"] != "" && u["name"] != "" && u["phone"] != "" && u["identity"]!=""  {

		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true) //连接数据库

		c := session.DB("test").C("people")

		var user User
		c.Find(bson.M{"name": u["name"]}).One(&user)
		if  user.Name ==""  {
			c.Insert(&u)
			i = 0 //数据库中之前不存在这个name，可以注册
		} else {
			i = 1 //数据库中已有这个name
		}
	} else {
		i = 2 //数据不完整
	}
	return i
}

func DeleteMember(u map[string]string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //连接数据库

	c := session.DB("test").C("people")

	c.Remove(bson.M{"name": u["name"]})
}

func ShowGroup(u map[string]string) []User {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //连接数据库

	c := session.DB("test").C("people")

	var users []User //用切片来存放所有查询结果
	c.Find(bson.M{"group": u["group"]}).All(&users)
	return users
}

func GetInformation(u map[string]string) []User {
	var users []User //用切片来存放所有查询结果
	if u["information"] == "yes" {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("test").C("people")


		c.Find(bson.M{"identity": "member"}).All(&users)
	}
	return users
}

func AddGroup(u map[string]string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	selector := bson.M{"name": u["name"]}
	data := bson.M{"$set": bson.M{"group": u["group"]}}
	c.Update(selector, data)
}

func ChangeInformation(u map[string]string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	selector := bson.M{"name": u["name"]}
	data := bson.M{
		"id":       u["id"],
		"password": u["password"],
		"email":    u["email"],
		"phone":    u["phone"],
		"name":     u["name"],
		"group":    u["group"],
		"identity": u["identity"],
		"status":   u["status"],
	}
	c.Update(selector, data)
}