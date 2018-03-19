package api

import (
	"blog/controllers"
	"blog/models"
)

type ApiarticleController struct {
	controllers.ToolsController
}

func (self *ApiarticleController) GetArticle() {
	self.QajaxList("成功", 1, 1, 1)
}

func (self *ApiarticleController) GetTagList(){
	tagList := models.TagGetList()
	self.QajaxList("成功", 0, 0, tagList)
}