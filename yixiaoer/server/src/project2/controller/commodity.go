package controller

import (
	"github.com/labstack/echo"
	"project/project2/model"
	"net/http"
	"fmt"
)

func ShowCategory(c echo.Context) error {
	CommodityCategory := map[string]string{
		"类别": "",
	}
	c.Bind(&CommodityCategory)
	fmt.Println(CommodityCategory)
	var commodity []model.Commodity
	commodity = model.ShowCategory(CommodityCategory)
	u := &commodity
	fmt.Println(commodity)
	return c.JSON(http.StatusOK, u)
}

func ShowLocation(c echo.Context) error {
	CommodityCategory := map[string]string{
		"地域": "",
	}
	c.Bind(&CommodityCategory)
	var commodity []model.Commodity
	commodity = model.ShowLocation(CommodityCategory)
	u := &commodity
	return c.JSON(http.StatusOK, u)
}

func CommodityInfo(c echo.Context) error {
	commodityInfo := map[string]string{
		"id": "",
		//"图片":"",
	}
	c.Bind(&commodityInfo)

	var commodity model.Commodity
	commodity = model.CommodityInfo(commodityInfo)
	u := &commodity
	fmt.Println(commodity)
	return c.JSON(http.StatusOK, u)
}

func PopluarRank(c echo.Context) error{
	userInfo := map[string]string{
		"hits":  "yes",
	}
	c.Bind(&userInfo)

	var commodity []model.Commodity

	commodity=model.PopularRank(userInfo)
	u := &commodity
	return c.JSON(http.StatusOK, u)
}


