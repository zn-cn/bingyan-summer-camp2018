package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

type User struct {
	Id       string
	Password string
	Email    string
	Phone    string
	Name     string
	Group    string
	Identity string
}

func UserSignup(loginuser User, user User) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	if err := c.Find(bson.M{"Id": loginuser.Id, "Password": loginuser.Password}).One(&user); err != nil {
		log.Fatal(err)
	}

}
