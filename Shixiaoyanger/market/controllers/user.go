package controllers

import (
	"fmt"
	"encoding/json"
	"market/models"
	"github.com/astaxie/beego"

	
)

type UserController struct {
	MainController
}

//注册：写入注册信息//1
func (this *UserController) Register(){	

	user := models.User{}
	err :=json.Unmarshal(this.Ctx.Input.RequestBody,&user)

	if err != nil {
		beego.Error(err)
		this.RetError(errInputData)
	}  
	 
	   ok,err1 := models.AddUser(&user)
	   
	if ok != true{
		if err1 != nil {
			beego.Error(err)
			this.RetError(errDatabase1)			
		} else{
			beego.Error(err)
			this.RetError(errDupUser)
		}
	}else{
		fmt.Println("jjjjjj",user.Username,"hhhhhhh")
		userstruct := UserStruct{
			User: user,
			StatusCode: sucregist, 
   		}
		this.Data["json"] = userstruct
    	this.ServeJSON()	
	}
	
	this.TplName = "review.html"
}


//登录验证.
func (this *UserController) Login(){
	var user models.User
	err :=json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		beego.Error(err)
	}

	u := user.Username
	p := models.GetDes(user.Password)

	ok,userInfo := models.FindbyUserame(u)
	if ok {
		if p == userInfo.Password {
			this.Data["json"] ="登陆成功"
			//设置session	
			this.SetSession("username",user.Username)
			}else{
				this.Data["json"] ="密码错误"
		}
	}else{
			this.Data["json"] ="账号不存在"	
		}
	
	this.ServeJSON()
	this.SetSession("username",user.Username)

	this.TplName = "info.html"
}


//退出注销.
func (this *UserController) Logout(){

	this.DelSession("username")

	this.TplName = "regist.html"
}

//
func (this *UserController) Update(){
	user := models.User{}
	/*
	fmt.Println("jjjjjj",user.Username,"hhhddhhhh")
	v := this.GetSession("username")
	fmt.Println("jjjdjjj",user.Username,"hhhhhhh")
	if v!= user.Username{
		this.Data["json"] = "userinfo incorrect"
		this.ServeJSON()
		return
	}
	*/
	err := json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	fmt.Println("jjjjjj",user.Username,"hhhhhhh")
	if err != nil{
		beego.Error(err)
	}
	err = models.UpdateUser(&user)
	if err !=nil{
		beego.Error(err)
	}

	_,user = models.FindbyUserame(user.Username)
	models.IncreaseView(&user)
	this.Data["json"] = "update success"
	this.ServeJSON()



	

	this.TplName = "regist.html"
}



