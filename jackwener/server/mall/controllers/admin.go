package controllers

import (
	"mall/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func (c *MainController) AddGoods() {
	var ob models.GoodJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	title := ob.Good[0].Title
	kind := ob.Good[0].Kind
	price := ob.Good[0].Price
	local := ob.Good[0].Local
	intro := ob.Good[0].Intro
	var good models.Good
	good.Intro = intro
	good.Price = price
	good.Kind = kind
	good.Local = local
	good.Id = title
	good.Views = 0
	o := orm.NewOrm()
	_, err := o.Insert(&good)
	var reJson models.GoodJson
	if err == nil {
		reJson.Status = 200
		reJson.Message = "添加成功"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	} else {
		reJson.Status = 400
		reJson.Message = "添加失败"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
}

func (c *MainController) DeleteGoods() {
	var ob models.GoodJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	title := ob.Good[0].Title
	o := orm.NewOrm()
	_, err := o.QueryTable("good").Filter("title",title).Delete()
	var reJson models.GoodJson
	if err == nil {
		reJson.Status = 200
		reJson.Message = "删除成功"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	} else {
		reJson.Status = 400
		reJson.Message = "删除失败"
		c.Data["json"] = reJson
		c.ServeJSON()
		return
	}
}

func (c *MainController) UpdateGoods() {
	var ob models.GoodJson
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Print(ob)
	good := ob.Good[0]
	o := orm.NewOrm()
	if good.Local != ""{
		o.QueryTable("good").Filter("title", good.Title).Update(orm.Params{"local": good.Local})
	}
	if good.Intro != ""{
		o.QueryTable("good").Filter("title", good.Title).Update(orm.Params{"intro": good.Intro})
	}
	if good.Kind != ""{
		o.QueryTable("good").Filter("title", good.Title).Update(orm.Params{"kind": good.Kind})
	}
	if good.Price != 0{
		var old models.Good
		o.QueryTable("good").Filter("title", good.Title).One(&old)
		if good.Price < old.Price {
			var cars []*models.Car
			_, err := o.QueryTable("car").Filter("Goods__Good__Title", good.Title).All(&cars)
			if err == nil {
				var name []string
				for i, _ := range cars{
					name[i] = cars[i].Id
				}
				var inform models.Inform
				for _, name := range name{
					inform.Id = name
					var user models.User
					o.QueryTable("user").Filter("name",name).One(&user)
					inform.User = &user
					inform.Content = "您收藏的商品"+name+"已经降价"
					o.Insert(&inform)
				}
			}
		}
		o.QueryTable("good").Filter("title", good.Title).Update(orm.Params{"price": good.Price})
	}
	var reJson models.GoodJson
	reJson.Status = 200
	reJson.Message = "修改成功"
	c.Data["json"] = reJson
	c.ServeJSON()
	return
}