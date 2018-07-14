package controller

import (
	"project/model"
	"net/http"
	"time"
	"github.com/labstack/echo"
)

type FormUser struct {
	Id       string `json:"id" form:"id"query:"id"`
	Password string `json:"password" form:"password"query:"password"`
	Email    string `json:"email" form:"email" query:"email"`
	Phone    string `json:"phone" form:"phone" query:"phone"`
	Name     string `json:"name" form:"name" query:"name"`
	Group    string `json:"group" form:"group" query:"group"`
	Identity string `json:"identity" form:"identity" query:"identity"`
}

func SearchId(user model.User) {
	return user.Id
}

func Cookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = SearchId(user)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func Request(c echo.Context) (err error) {
	u := new(FormUser)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func ResponseJsonyes(c echo.Context) error {
	u := &model.User{
		status: true
	}
	return c.JSON(http.StatusOK, u)
}

func ResponseJsonno(c echo.Context) error {
	u := &model.User{
		status: false
	}
	return c.JSON(http.StatusOK, u)
}

func Login(user model.User) {
	if user != nil {
		Request(c)
		Cookie(c)
		ResponseJsonyes(c)
	} else {
		ResponseJsonno(c)
	}
}
