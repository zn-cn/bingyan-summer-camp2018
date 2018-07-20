package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       string `json:"id" form:"id"query:"id"`
	Password string `json:"password" form:"password"query:"password"`
	Email    string `json:"email" form:"email" query:"email"`
	Phone    string `json:"phone" form:"phone" query:"phone"`
	Name     string `json:"name" form:"name" query:"name"`
	Hits     string `json:"hits" form:"hits" query:"hits"`
}

//var Publickey = `-----BEGIN 公钥-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDS3pFbkT33ioFIh7kFb/zByVnU
//lQ5OnFvSLYXM42LmzS9ppE5nvk5VBLF1MiwQlHCrM/94cxVQ+3jBZPkrmx8Ulii8
//9kAXn+YGgpApUYiC/5tRdSY8C4TkqHbJD/wEgweTsD3CRBxSMyfKKRUh4/Wu7Trx
//FPiBDvurxCqQyC/vpQIDAQAB
//-----END 公钥-----
//`
//
//var Pirvatekey = `-----BEGIN 私钥-----
//MIICoTAbBgkqhkiG9w0BBQMwDgQI4TCFXuyfWGUCAggABIICgNemXWIGNtUETmtR
//vP5QkQ+ItX7DtsX/K3NHLiEaG0IDyoeiBX9nwt9TbQk/erCTBbp4N/YOXYVlMcws
//FTPhaCIEJqYnM/ZYhLHz0gofEQaFqTvJYe04mF7kY2B1UujjinZhEJMpmuf4kxOw
//NdD9zCU5/YC00qZlTfdXo11z5e5SF/8bBTCzFnsgoI16SmhHq2NbMsZoRDS7EhDb
//h+t3UWS/QGk8t5DBc63UTlzxP0yA6Ef1/eMRkgAOfJPP9ED7EHxP81zBLj/U6pC8
//06lvK/qJOMD+WtwLdP4mNAvyoHAjUvoROtXxTldnoBUaRoO853EuCvzhbm0str63
//y4khPflzLsPzB3MW1hftjS2cPNjliPntBMdk+sPhDbXAtwrzMn1/4oTtvJU/kKN/
//Iz22o7a6dtScc9PanS/RIr6AtZBWbP46+dv66yx+4J66Z8V2TnmzVYJIQLt3Y5cj
//ai7/d0X6WKr1XN0eqFxDwniUgrdNzK9SCi4kd0XypTg9lNQzqF/V9bO5f3oDJzYS
//2eH/PMgtLGa6H3d7DezQzw89MELVuWvaG8UdsGhhjsAtAWfgOFo2KRBpGAbIZbIp
//uAXIIUTb6JZT3vozVOsGXOYUEK9J8FNovTUGAXNlZhC5xlyRlezHRd2dKLRpmbHq
//LDsJJaN6g1q4u1mKBPlZK3oEdAv6kODdkcKaJbbwyAE3KfzOpFrczeequ/5mVLrY
//8edC4l+IWDoxE/QmgLWkbSkjgQEXnPw3p6xCfyVgSMdDHAn2S2nZm8offeUMZj8R
//pbwPL3evdevZzaADoABv9OejCpep2nK1/sFMzKUYLxSY3+00/74872Gy49PHbfDi
//0APM9yg=
//-----END 私钥-----
//`

func Login(u map[string]string) int8 {
	var i int8
	if u ["name"] != "" && u["password"] != "" {

		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true) //连接数据库

		c := session.DB("商城").C("用户")

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
	if u["password"] != "" && u["name"] != "" && u["phone"] != "" {

		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true) //连接数据库

		c := session.DB("商城").C("用户")

		var user User
		c.Find(bson.M{"name": u["name"]}).One(&user)
		if user.Name == "" {
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

func UserInfo(u map[string]string)User{
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //连接数据库

	c := session.DB("商城").C("用户")

	var user User
	c.Find(bson.M{"name": u["name"]}).One(&user)
	return user
}

func UserHits(u map[string]string){
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("商城").C("用户")

	selector := bson.M{"name": u["name"]}
	data := bson.M{"$set": bson.M{"hits": u["hits"]}}
	c.Update(selector, data)
}