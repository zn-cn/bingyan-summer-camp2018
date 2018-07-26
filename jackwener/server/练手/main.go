package main

import "github.com/astaxie/beego"

import (
	 "hello/models"
	"hello/routers"
)

func main() {
	models.Init()
	routers.AdminInit()
	routers.IndexInit()
	beego.Run()
}