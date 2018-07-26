package controllers

import (
	"github.com/astaxie/beego"
	"market/models"
)
// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

// ControllerError is controller error info structer.
type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}
type UserStruct struct{
	User 	 	 models.User
	StatusCode   *StatusCode
}
type GoodsStruct struct{
	Goodsinfo 	[]*models.Goods
	StatusCode  *StatusCode
}
type StatusCode struct{
	Status   int64   	 `json:"status"`
	Code     int64   	 `json:"code"`
	Message  string 	 `json:"message"`
	DevInfo  string 	 `json:"dev_info"`
	MoreInfo string 	 `json:"more_info"`

}

// Predefined controller error values.
var (
	err404          = &ControllerError{404, 404, "page not found", "page not found", ""}
	errInputData    = &ControllerError{400, 10001, "数据输入错误", "客户端参数错误", ""}
	errDatabase1    = &ControllerError{500, 10002, "服务器错误", "数据库操作错误", ""}
	errDupUser      = &ControllerError{400, 10003, "用户信息已存在", "数据库记录重复", ""}
	errNoUser       = &ControllerError{400, 10004, "用户信息不存在", "数据库记录不存在", ""}
	errPass         = &ControllerError{400, 10005, "用户信息不存在或密码不正确", "密码不正确", ""}
	errNoUserPass   = &ControllerError{400, 10006, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确", ""}
	errNoUserChange = &ControllerError{400, 10007, "用户信息不存在或数据未改变", "数据库记录不存在或数据未改变", ""}
	errInvalidUser  = &ControllerError{400, 10008, "用户信息不正确", "Session信息不正确", ""}
	errOpenFile     = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
	errWriteFile    = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
	errSystem       = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
	errExpired      = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
	errPermission   = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}

	errNoGoods       = &ControllerError{400, 20001, "商品信息不存在", "数据库记录不存在", ""}
	errDatabase2     = &ControllerError{500, 20002, "服务器错误", "数据库操作错误", ""}
)
//Predefined Statuscode values.
var (
	sucregist      = &StatusCode{200,2001,"注册成功","",""}
	sucgoodsinfo      =&StatusCode{200,2002,"查询成功","",""}
	sucgosearch      =&StatusCode{200,2003,"","",""}
	suc2      =&StatusCode{200,2004,"","",""}
	suc3      =&StatusCode{200,2005,"","",""}
	suc4      =&StatusCode{200,2006,"","",""}
	suc5      =&StatusCode{200,2007,"","",""}
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Tel"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
//return error info
func (this *MainController)  RetError(e *ControllerError){
	this.Data["json"] = e
	this.ServeJSON()

}
