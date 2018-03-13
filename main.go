package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"blog/models"
)

func main() {
	models.Init()
	beego.Run()
}

