package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"blog/models"
	"blog/libs"
	"mime/multipart"
	"time"
	"os"
)

type BaseController struct{
	ToolsController
	controllerName  string
	actionName      string
	user 			*models.Admin
	userId			int
	userName	    string
	loginName	    string
	pageSize		int
	allowUrl        string
}

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

//初始化
func (self *BaseController) Prepare(){
	controllerName, actionName := self.GetControllerAndAction()
	// -10 为了去掉controller
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName

	if(strings.Compare(self.controllerName, "apidoc")) != 0 {
		self.auth()
	}
	self.Data["loginUserName"] = self.userName
}

//登录权限验证
func (self *BaseController) auth()  {
	//获取cookie里的值
	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2{
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if(userId > 0){
			user, err := models.AdminGetById(userId)
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)){
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
			}
			//验证权限
			//isHasAuth := strings.Contains(self.allowUrl, self.controllerName + "/" + self.actionName)
			//noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			//isNoAuth := strings.Contains(noAuth, self.actionName)
			//if isHasAuth == false && isNoAuth == false{
			//	self.Ctx.WriteString("没有权限")
			//	self.ajaxMsg("没有权限", MSG_ERR)
			//	return
			//}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin"){
		self.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

//用户权限，拼接左侧菜单
func (self *BaseController) AdminAuth()  {
	//左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if self.userId != 1{
		adminAuthIds, _ := models.RoleAuthGetByIds(self.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _,v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/"{
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}
	self.Data["SideMenu1"] = list[:i]  //一级菜单
	self.Data["SideMenu2"] = list2[:j] //二级菜单
	self.allowUrl = allow_url + "/home/index"
}


func (self *BaseController) display(tpl ...string)  {
	var tplname string
	if len(tpl)>0{
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	}else{
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

//上传图片,返回值为原图及缩略图路径
func (self *BaseController) SaveFile(f multipart.File,h *multipart.FileHeader,par_name string)(err error,l string,m string,s string){
	defer f.Close()
	now := time.Now()
	upload_path := beego.AppConfig.String("uploadpath")
	dir := upload_path+ strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	dirf := strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())+ "/"
	err = os.MkdirAll(dir, 0755)

	if err != nil {
		return
	}
	filename := h.Filename
	newname := libs.FormatFileName(filename)
	imgpath := dir+"/"+newname
	err = self.SaveToFile(par_name, imgpath)

	if err != nil{
		return
	}else {
		image,_ := beego.AppConfig.GetSection("image")
		width_m,_ := strconv.Atoi(image["width_m"])
		width_s,_ := strconv.Atoi(image["width_s"])
		height_m,_ := strconv.Atoi(image["height_m"])
		height_s,_ := strconv.Atoi(image["height_s"])
		err = libs.MakeMilThumb(imgpath,width_m,height_m)
		err = libs.MakeSmallThumb(imgpath,width_s,height_s)
		if err != nil {
			return
		}else {
			l = dirf+newname
			m = dirf+strings.Replace(newname, "_l.", "_m.",1)
			s = dirf+strings.Replace(newname, "_l.", "_s.",1)
			return err,l,m,s
		}
	}

}
