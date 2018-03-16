package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type ToolsController struct {
	beego.Controller
}

// 是否POST提交
func (self *ToolsController) isPost() bool{
	return self.Ctx.Request.Method == "POST"
}

// 获取用户IP地址
func (self *ToolsController) getClientIp() string{
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self * ToolsController) redirect(url string){
	self.Redirect(url, 302)
	self.StopRun()
}


func (self *ToolsController) ajaxMsg(msg interface{}, msgno int)  {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表
func (self *ToolsController) ajaxList(msg interface{}, msgno int, count int64, data interface{})  {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

