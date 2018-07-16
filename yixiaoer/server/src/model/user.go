package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
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
	Status   string
}

func Login(u map[string]string) bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	var user User
	if err := c.Find(bson.M{"Id": u["id"], "Password": u["password"]}).One(&user); err != nil {
		log.Fatal(err)
		return true
	}
	return false
}

func SignUp(u map[string]string) bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	var user User
	if err := c.Find(bson.M{"Id": u["id"]}).One(&user); err == nil {
		err = c.Insert(&u)
		if err != nil {
			log.Fatal(err)

		}
		return true
	}
	return false
}

func DeleteMember(u map[string]string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	err = c.Remove(bson.M{"id": u["id"]})
}

func ShowGroup(u map[string]string) []User {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	var users []User //用切片来存放所有查询结果
	c.Find(bson.M{"group": u["group"]}).All(&users)
	return users
}

func GetInformation(user []User) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	c.Find(bson.M{"identity": "member"}).All(&user)
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

	selector := bson.M{"id": u["id"]}
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

	selector := bson.M{"id": u["id"]}
	data := bson.M{
		"id":       u["id"],
		"password": u["password"],
		"email":    u["email"],
		"phone":    u["phone"],
		"name":     u["name"],
		"group":    u["group"],
		"identity": u["identity"],
		"status":   "1",
	}
	c.Update(selector, data)
}
