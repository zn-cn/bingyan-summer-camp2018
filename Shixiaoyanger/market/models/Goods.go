package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"



)
type Goodses []Goods

type Goods struct{
	Id   		 int64   //	`json:"id"`
	Goodsname 	 string		`json:"goodsname"`	
	Category 	 string 	`json:"category"`
	Position 	 string		`json:"position"`
	Introduction string		`json:"introduction"`
	Title        string		`json:"title"`
	Price        int64		`json:"price"`
	Goodviews    int64		`json:"goodviews"`
	Viewtime     int64		`json:"viewtime"`
}

//locate goods by goodsname
func Readgoods(goods *Goods) error{
	o := orm.NewOrm()
	qs := o.QueryTable("goods")
	err := qs.Filter("goodsname",goods.Goodsname).One(&goods)
	return err
	//err == nil goodsname has existed
	//err!= nil  .....not exist

}

//添加商品
func Addgoods(goods *Goods)  (bool, error){
	o := orm.NewOrm()
	err := Readgoods(goods)
	if err == nil{
		return false, err //goodsname has existed
	}
	_,err =o.Insert(goods)
	if err != nil{
		return false, err  //faild to insert
	}
	return true, nil
}
//删除商品
func Deletegoods(goods *Goods) error {
	o := orm.NewOrm()
	_,err := o.Delete(goods)
	if err!= nil{
		return err
	}
	return nil
}

func Updategoods(goods *Goods) error{
	o := orm.NewOrm()
	err := Readgoods(goods)
	if err != nil{
		return err//not exist
	}
	_,err =o.Update(goods,"views","viewtime")
	return err //update success
}
//浏览量+1
func IncreGoodsViews(goods *Goods) {
	o := orm.NewOrm()
	goods.Goodviews += 1
	o.Update(goods,"goodviews")
}
func UpadateViewtime(goods *Goods){
	o := orm.NewOrm()

		 // 获取当前(当地)时间
		 t := time.Now()
		 // 获取0时区时间
		 t = time.Now().UTC()
		 fmt.Println(t)
		 // 获取当前时间戳
		 timestamp := t.Unix()
		 fmt.Println(timestamp)
		 // 获取时区信息
		 name, offset := t.Zone()
		 fmt.Println(name, offset)
		 // 把时间戳转换为时间
		 currenttime := time.Unix(timestamp+int64(offset), 0)
		 // 格式化时间
		 fmt.Println("Current time : ", currenttime.Format("2006-01-02 15:04:05"))

	goods.Viewtime = timestamp
	o.Update(goods,"viewtime")

}

/********************查找商品**************************/
func SearchByCategory(category string) (bool, []*Goods){
	o :=orm.NewOrm()
	var goods []*Goods
	fmt.Println("123242423")
	_,err := o.QueryTable("goods").Filter("category", category).All(&goods)//All(&goods,"","")
	for _,value := range goods{
		IncreGoodsViews(value)
		fmt.Println(value)
	}
	fmt.Println("eiwrh98233h9wn")
	return err!= orm.ErrNoRows, goods
}

func SearchByPosition(position string) (bool, []*Goods){
	o :=orm.NewOrm()
	var goods []*Goods

	num,err := o.QueryTable("goods").Filter("position", position).All(&goods)
	fmt.Println("返回数据条数",num,)
	for _,value := range goods{
		IncreGoodsViews(value)
		UpadateViewtime(value)
		fmt.Println(value)
	}
	return err!= orm.ErrNoRows, goods
}

func SearchById(id string) (bool, Goods){
	o :=orm.NewOrm()
	var goods Goods
	err := o.QueryTable("goods").Filter("id", id).One(&goods)
	IncreGoodsViews(&goods)
	UpadateViewtime(&goods)
	return err!= orm.ErrNoRows, goods
}

/********************热门查询、最新查询 *******************************/
//热门查询
func PopularSearch() (error,[]*Goods){
	var goods []*Goods
	 _,err:= orm.NewOrm().Raw("SELECT * from goods order by goodviews desc limit 4" ).QueryRows(&goods)
	 if err != nil{
		 beego.Error(err)
		 fmt.Println("失败")
	 }
	fmt.Print("GJJH",goods)

	return err, goods
}
//最新查询
func LatestSearch() (error, []*Goods){
	var goods []*Goods
	 _,err:= orm.NewOrm().Raw("SELECT * from goods order by viewtime desc limit 4" ).QueryRows(&goods)
	 if err != nil{
		 beego.Error(err)
		 fmt.Println("失败")
	 }
	fmt.Print("dsfsdfs",goods)

	return err, goods

}

