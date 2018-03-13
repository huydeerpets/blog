package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	BaseController
}

func (self *UploadController) UploadImgOne()  {

	f,h,err_img := self.GetFile("file")
	if err_img != nil{
		self.ajaxMsg(err_img.Error(), MSG_ERR)
	}
	//定义原图,中图,大图路径
	//定义图片的name属性
	var par_name = "file"
	err,filepath,filepath_m,filepath_s := self.SaveFile(f,h,par_name)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	list := make(map[string]interface{})
	list["filepath"] = filepath
	list["filepath_m"] = filepath_m
	list["filepath_s"] = filepath_s
	self.ajaxList("成功", MSG_OK, 1, list)
}

func (self *UploadController) UploadImgEditormd()  {

	f,h,err_img := self.GetFile("editormd-image-file")
	if err_img != nil{
		out := make(map[string]interface{})
		out["success"] = MSG_ERR
		out["message"] = err_img.Error()
		self.Data["json"] = out
		self.ServeJSON()
		self.StopRun()
	}
	//定义原图,中图,大图路径
	//定义图片的name属性
	var par_name = "editormd-image-file"
	err,filepath, _, _ := self.SaveFile(f,h,par_name)
	if err != nil{
		out := make(map[string]interface{})
		out["success"] = MSG_ERR
		out["message"] = err_img.Error()
		self.Data["json"] = out
		self.ServeJSON()
		self.StopRun()
	}
	out := make(map[string]interface{})
	out["success"] = 1
	out["message"] = "成功"
	out["url"] = beego.AppConfig.String("site.url") + "/" + beego.AppConfig.String("uploadpath") + filepath
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

