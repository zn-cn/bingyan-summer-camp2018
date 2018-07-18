package main

import "github.com/astaxie/beego"

import (
	_ "hello/models"
	_ "hello/routers"
)

func main() {
	beego.Run()
}