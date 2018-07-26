package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Commodity struct {
	Id       bson.ObjectId ` json:"id" form:"id" query:"id" bson:"_id"`
	Title    string        `json:"title" form:"title" query:"title" bson:"title"`
	Info     string        `json:"info" form:"info" query:"info" bson:"info"`
	Price    int           `json:"price" form:"price" query:"price" bson:"price"`
	Picture  string        `json:"picture" form:"picture" query:"picture" bson:"picture"`
	Category string        `json:"category" form:"category" query:"category" bson:"category"`
	Location string        `json:"location" form:"location" query:"location" bson:"location"`
	Hits     int           `json:"hits" form:"hits" query:"hits" bson:"hits"`
}

//按类别查询
func ShowCategory(u map[string]string) []Commodity {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //连接数据库

	c := session.DB("商城").C("商品")

	var commodities []Commodity//用切片来存放所有查询结果
	err = c.Find(bson.M{"category": u["category"]}).All(&commodities)
	fmt.Println("test")
	fmt.Println(err)
	fmt.Println(commodities)
	fmt.Println("test over")
	return commodities
}

//按地域查询
func ShowLocation(u map[string]string) []Commodity {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //连接数据库

	c := session.DB("商城").C("商品")

	//var commodities []Commodity //用切片来存放所有查询结果
	var commodities []Commodity
	c.Find(bson.M{"location": u["location"]}).All(&commodities)
	return commodities
}

//商品页面

func CommodityInfo(u map[string]string) Commodity {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("商城").C("商品")

	var commodity Commodity
	c.Find(bson.M{"_id": bson.ObjectIdHex(u["id"])}).One(&commodity)

	//c.Find(bson.M{"id": u["id"]}).One(&commodity)
	//c.Find(bson.M{"图片": u["图片"]}).One(&commodity)

	return commodity
}

//热度查看
func PopularRank(u map[string]string) []Commodity {
	var commodities []Commodity
	if u["hits"] == "yes" {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("商城").C("商品")

		c.Find(nil).Sort("-hits").All(&commodities) // 按照点击量升序排列
	}
	return commodities
}

func CommodityHits(u map[string]string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("商城").C("商品")

	selector := bson.M{"title": u["title"]}
	data := bson.M{"$set": bson.M{"hits": u["hits"]}}
	c.Update(selector, data)
}