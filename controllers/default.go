package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//首页
func (self *MainController) Index() {
	self.TplName = "default/index.html"
}

//文章列表
func (self *MainController) Detail(){
	self.TplName = "default/article.html"
}

//时间墙
func (self *MainController) Timeline(){
	self.TplName = "default/timeline.html"
}

//关于
func (self *MainController) About(){
	self.TplName = "default/about.html"
}

