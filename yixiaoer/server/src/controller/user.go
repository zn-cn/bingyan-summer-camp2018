package controller

import (
	"project/model"
	"net/http"
	"time"
	"github.com/labstack/echo"
)

func Request(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func Login(c echo.Context) error {
	Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
	}
	c.Bind(&userInfo)
	if model.Login(userInfo) == true {
		//userInfo  := map[string]string{
		//	"id": "",
		//	"pw": "",
		//}
		//c.Bind(&userInfo)
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = userInfo["id"]
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		u := &model.User{
			Name: "yes",
		}
		return c.JSON(http.StatusOK, u)
	} else {
		u := &model.User{
			Name: "no",
		}
		return c.JSON(http.StatusOK, u)
	}
}

func SignUp(c echo.Context) error {
	Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "0",
	}
	c.Bind(&userInfo)
	if model.SignUp(userInfo) == true {
		u := &model.User{
			Name: "yes",
		}
		return c.JSON(http.StatusOK, u)
	} else {
		u := &model.User{
			Name: "no",
		}
		return c.JSON(http.StatusOK, u)
	}
}

func DeleteMember(c echo.Context) error {
	Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "1",
	}
	c.Bind(&userInfo)
	model.DeleteMember(userInfo)
	return c.NoContent(http.StatusOK)
}

func ShowGroup(c echo.Context) error {
	Request(c)
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
	Request(c)
	var member []model.User
	model.GetInformation(member)
	u := &member
	return c.JSON(http.StatusOK, u)
}

func AddGroup(c echo.Context) error {
	Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "1",
	}
	c.Bind(&userInfo)
	model.AddGroup(userInfo)
	return c.NoContent(http.StatusOK)
}

func ChangeInformation(c echo.Context) error {
	Request(c)
	userInfo := map[string]string{
		"id":       "",
		"password": "",
		"email":    "",
		"phone":    "",
		"name":     "",
		"group":    "",
		"identity": "",
		"status":   "1",
	}
	c.Bind(&userInfo)
	model.ChangeInformation(userInfo)
	return c.NoContent(http.StatusOK)
}
