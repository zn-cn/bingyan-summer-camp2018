package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"project/project2/model"
	"time"
	"crypto/md5"
	"encoding/hex"
)

func Login(c echo.Context) error {
	//Request(c)
	userInfo := map[string]string{
		"name":     "",
		"password": "",
	}
	c.Bind(&userInfo)

	h := md5.New()
	h.Write([]byte(userInfo["password"])) // 需要加密的字符串为密码
	has := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
	userInfo["password"] = has

	var u map[string]string
	if model.Login(userInfo) == 0 { //密码与账户匹配
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = userInfo["name"]
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		u = map[string]string{
			"status": "yes",
		}
	} else if model.Login(userInfo) == 1 { //有name但是pw不匹配
		u = map[string]string{
			"status": "wrong pw",
		}
	} else if model.Login(userInfo) == 2 { //没有name
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
	}
	c.Bind(&userInfo)
	//data := []byte(userInfo["password"])
	//has := md5.Sum(data)
	h := md5.New()
	h.Write([]byte(userInfo["password"])) // 需要加密的字符串为密码
	has := hex.EncodeToString(h.Sum(nil)) // 输出加密结果
	userInfo["password"] = has
	var u map[string]string
	if model.SignUp(userInfo) == 0 {
		u = map[string]string{
			"status": "yes",
		}
	} else if model.SignUp(userInfo) == 1 {
		u = map[string]string{
			"status": "already have",
		}
	} else if model.SignUp(userInfo) == 2 {
		u = map[string]string{
			"status": "incomplete data",
		}
	}
	return c.JSON(http.StatusOK, u)
}

func UserInfo(c echo.Context)error {
	userInfo := map[string]string{
		"name":     "",
	}
	c.Bind(&userInfo)
	model.UserHits(userInfo)
	var user model.User
	user = model.UserInfo(userInfo)
	u := &user
	return c.JSON(http.StatusOK, u)
}