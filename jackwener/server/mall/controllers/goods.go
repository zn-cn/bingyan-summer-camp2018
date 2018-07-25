package controllers

import (
	"mall/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"fmt"
	"os"
)

// 种类查询
func (c *MainController) KindSearch() {
	var ob models.GoodJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	kind := ob.Good[0].Kind
	o := orm.NewOrm()
	var goods []models.Good
	num, err := o.QueryTable("good").Filter("kind", kind).All(&goods)
	var reJson models.GoodJson
	if err != nil{
		reJson.Status = 404
		reJson.Message = "查询失败"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	} else if num == 0 {
		reJson.Status = 204
		reJson.Message = "分类为空"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
	var record models.Record
	err1 := o.QueryTable("record").Filter("content",kind).One(&record)
	if err1 == nil {
		o.QueryTable("record").Filter("content", kind).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, 1)})
	} else {
		record.Content = kind
		record.Number = 1
		o.Insert(&record)
	}
	reJson.Good = goods
	reJson.Status = 200
	reJson.Message = "查询成功"
	c.Data["json"] = reJson
	c.ServeJSON()
}

// 地域查询
func (c *MainController) LocalSearch() {
	var ob models.GoodJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	local := ob.Good[0].Local
	o := orm.NewOrm()
	var goods []models.Good
	num, err := o.QueryTable("good").Filter("local", local).All(&goods)
	var reJson models.GoodJson
	if err != nil{
		reJson.Status = 404
		reJson.Message = "查询失败"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	} else if num == 0 {
		reJson.Status = 204
		reJson.Message = "分类为空"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
	var record models.Record
	err1 := o.QueryTable("record").Filter("content",local).One(&record)
	if err1 == nil {
		o.QueryTable("record").Filter("content", local).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, 1)})
	} else {
		record.Content = local
		record.Number = 1
		o.Insert(&record)
	}
	reJson.Good = goods
	reJson.Status = 200
	reJson.Message = "查询成功"
	c.Data["json"] = reJson
	c.ServeJSON()
}

// 商品信息页
func (c *MainController) Goods() {
	title := c.GetString(":good")
	fmt.Print(title)
	var good models.Good
	o := orm.NewOrm()
	err := o.QueryTable("good").Filter("title", title).One(&good)
	fmt.Print(good)
	if err != nil {
		var reJson models.GoodJson
		reJson.Message = "商品不存在"
		reJson.Status = 404
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
	local := good.Local
	//fmt.Printf(local)
	intro := good.Intro
	price := good.Price
	url   := "static/img/" + title +"/"
	var reJson = models.GoodJson{}
	var good1 models.Good
	good1.Local = local
	good1.Price = price
	good1.Intro = intro
	reJson.Good = append(reJson.Good,good1)
	reJson.Url = url
	reJson.Status = 200
	reJson.Message = "返回成功"
	c.Data["json"] = reJson
	c.ServeJSON()
	return
}

// 上传图片
func (c *MainController) Picture() {
	name := c.GetString(":good")
	f, h, err := c.GetFile("name")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	path := "static/img/"+name
	_, err1 := os.Stat(path)
	if err1 != nil {
		err2 := os.Mkdir(path, os.ModePerm)
		if err2 != nil{
			fmt.Print(err2)
		}
	}
	c.SaveToFile("name", "static/img/"+name+"/"+ h.Filename) // 保存位置在 static/static/img/name/, 没有文件夹要先创建
}



// 最新查询
func (c *MainController) RecentSearch() {
	var reJson models.SearchJson
	var records  []models.Record
	o := orm.NewOrm()
	_, err := o.QueryTable("record").OrderBy("-date").All(&records)
	if err != nil{
		o.QueryTable("record").All(&records)
	}
	fmt.Print(records)
	for i, _ := range records{
		reJson.Content = append(reJson.Content,records[i].Content)
	}
	reJson.Message = "查询成功"
	reJson.Status = 200
	c.Data["json"] = reJson
	c.ServeJSON()
	return
}

// 热门查询
func (c *MainController) PopularSearch() {
	var reJson models.SearchJson
	var records  []models.Record
	o := orm.NewOrm()
	_, err := o.QueryTable("record").OrderBy("-number").All(&records)
	if err != nil{
		o.QueryTable("record").All(&records)
	}
	fmt.Print(records)
	for i, _ := range records{
		reJson.Content = append(reJson.Content,records[i].Content)
	}
	reJson.Message = "查询成功"
	reJson.Status = 200
	c.Data["json"] = reJson
	c.ServeJSON()
	return
}


/*
func (c *MainController) Car() {
	var ob models.ReJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	userId := ob.Users[0].Id
	o := orm.NewOrm()
	var goods[] class.Goods
	var carId int
	o.QueryTable("user").Filter("id", userId).RelatedSel().One(&carId,"carId")
	o.QueryTable("good").Filter("car_id",carId).All(&goods)
	var reJson models.ReJson
	reJson.Goods = goods
	reJson.Status = 200
	c.Data["json"] = reJson
	c.ServeJSON()
}

func (c *MainController) PopularSearch() {

}

func (c *MainController) PopularSearch() {

}
*/