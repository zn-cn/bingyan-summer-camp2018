package controllers

import (
	"fmt"

	"encoding/json"
	"market/models"
	"github.com/astaxie/beego"
)

type GoodsController struct{
	MainController
}
type SearchBy struct{
	Types      string    `json:"types"`  
	Details    string    `json:"details"`
	//按类别 // 1 电子设备
	//001	// 2 书籍资料
			// 3 宿舍百货
			// 4 美妆护肤
			// 5 女装
			// 6 男装
			// 7 鞋帽配饰
			// 8 门票卡券
			// 9 其他

	//按地域 // 1 韵苑
	//002	// 2 沁苑
			// 3 紫菘
			// 4 其他
}

//查询并返回商品信息
func (this *GoodsController)  Search(){
	var op SearchBy
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&op)
	if err != nil{
		beego.Error(err)
	}
	var goodsinfo []*models.Goods
	fmt.Println("dasa",op.Types,"jutyht")
	var ok bool

	if op.Types =="001"{

		ok, goodsinfo = models.SearchByCategory(op.Details)

	}
	if op.Types == "002"{

		ok, goodsinfo = models.SearchByPosition(op.Details)

	}
	if ok{
		goodsstruct :=GoodsStruct{
			Goodsinfo:  goodsinfo,
		}
		this.Data["json"] = goodsstruct 
		this.ServeJSON()

	}
}
//查询单个商品信息//1
func (this *GoodsController)  View(){

	id := this.Ctx.Input.Param(":id")
	ok,goods := models.SearchById(id)
	if ok{
		goodsstrct := GoodsStruct{
			Goodsinfo: []*models.Goods{&goods},
			StatusCode: sucgoodsinfo,
		}
		this.Data["json"] = goodsstrct
		this.ServeJSON()
	}
	
	this.RetError(errNoGoods)
	return
}
//添加商品信息
func (this *GoodsController)  Add(){
	var goinfo models.Goods
	err0 := json.Unmarshal(this.Ctx.Input.RequestBody,&goinfo)
	if err0 != nil{
		beego.Error(err0)
	}
	ok, err := models.Addgoods(&goinfo)
	if ok{
		this.Data["json"] = "添加成功"
	}else{
		if err == nil{
			this.Data["json"] = "商品名已存在"
		}else{
			this.Data["json"] = "添加失败"
		}
	}
    this.ServeJSON()
}


func (this *GoodsController)  Popular(){

	err,goods := models.PopularSearch()
	if err != nil{
		beego.Error(err)
		this.RetError(errDatabase2)
		return
	}
	goodsstrct := GoodsStruct{
		Goodsinfo:   goods ,                  //[]*models.Goods{&goods},
		StatusCode: sucgoodsinfo,
	}
	this.Data["json"] = goodsstrct
	this.ServeJSON()
	return
	
}
func (this *GoodsController)  Latest(){

	err,goods := models.LatestSearch()
	if err != nil{
		beego.Error(err)
		this.RetError(errDatabase2)
		return
	}
	goodsstrct := GoodsStruct{
		Goodsinfo:   goods ,                
		StatusCode: sucgoodsinfo,
	}
	this.Data["json"] = goodsstrct
	this.ServeJSON()
	return
	
}



func (this *GoodsController) Test(){
	this.TplName = "test.html"
}



/*time 操作

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
		
*/		
		
		
		
		
		
		
		
		
		