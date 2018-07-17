package controller

import (
	"project/model"
	"net/http"
	"time"
	"github.com/labstack/echo"
	"fmt"
)

//func Request(c echo.Context) (err error) {
//	u := new(model.User)
//	if err = c.Bind(u); err != nil {
//		return
//	}
//	return c.JSON(http.StatusOK, u)
//}

func Login(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		"name":     "",
		"password": "",
	}
	c.Bind(&userInfo)
	fmt.Println(userInfo)
	var u map[string]string
	if model.Login(userInfo) == 0 {   //密码与账户匹配
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = userInfo["name"]
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		u = map[string]string{
			"status": "yes",
		}
	} else if model.Login(userInfo) == 1 {  //有name但是pw不匹配
		u = map[string]string{
			 "status":"wrong pw",
		}
	} else if model.Login(userInfo)==2{               //没有name
		u = map[string]string{
			"status": "no user",
		}
	}
	return c.JSON(http.StatusOK, u)
}

func SignUp(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		//"id":     "",在数据库中会自动分配一个id所以在注册时可以不需要设置id
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "0", //此时成员的状态还未被验证
	}
	c.Bind(&userInfo)
	fmt.Println(userInfo)
	var u map[string]string
	if model.SignUp(userInfo) == 0 {
		u = map[string]string{
			"status": "yes",
		}
	} else if model.SignUp(userInfo) == 1 {
		u = map[string]string{
			"status": "already have",
		}
	} else if model.SignUp(userInfo) == 2{
		u = map[string]string{
			"status": "incomplete data",
		}
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteMember(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		//"id":       "",
		//"password": "",
		//"email":    "",
		//"phone":    "",
		  "name":     "",
		//"group":    "",
		//"identity": "",
		//"status":   "",
	}
	c.Bind(&userInfo)
	model.DeleteMember(userInfo)

		u := map[string]string{
			"status": "yes",
		}
	return c.JSON(http.StatusOK, u)
}

func ShowGroup(c echo.Context) error {
	//Request(c)
	userGroup := map[string]string{
		"group": "",
	}
	c.Bind(&userGroup)
	var user []model.User
	user = model.ShowGroup(userGroup)
	u := &user
	return c.JSON(http.StatusOK, u)
}

func GetInformation(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		"information": "",
	}
	c.Bind(&userInfo)
	var member []model.User
	member=model.GetInformation(userInfo)
	u :=&member
	return c.JSON(http.StatusOK, u)
}

func AddGroup(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		//"id":       "",
		//"password": "",
		//"email":    "",
		//"phone":    "",
		  "name":     "",
		  "group":    "",
		//"identity": "",
		  "status":   "1",
	}
	c.Bind(&userInfo)
	model.AddGroup(userInfo)
	return c.NoContent(http.StatusOK)
}

func ChangeInformation(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "",
	}
	c.Bind(&userInfo)
	model.ChangeInformation(userInfo)
	return c.NoContent(http.StatusOK)
}
