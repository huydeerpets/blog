package api

import (
	"blog/controllers"
	"blog/models"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ApiEmeiymCommentController struct {
	controllers.ToolsController
}

func (self *ApiEmeiymCommentController) PostComment() {

	name := strings.TrimSpace(self.GetString("name"))
	phone := strings.TrimSpace(self.GetString("phone"))
	content := strings.TrimSpace(self.GetString("content"))
	source, err := self.GetInt("source")
	//if err != nil {
	//	self.QajaxMsg(err.Error(), -1)
	//}

	if len(name) <= 0 || len(phone) <= 0 || len(content) <= 0 {
		self.QajaxMsg("请填写内容", -1)
	}

	evc := new(models.EmeiymVisitorComment)
	evc.Name = name
	evc.Phone = phone
	evc.Content = content
	evc.Source = source
	evc.CreateTime = time.Now().Unix()

	_, err = orm.NewOrm().Insert(evc)

	if err != nil {
		self.QajaxList(err.Error(), -1, 0, "")
	}

	self.QajaxList("成功", 0, 0, evc)
}
