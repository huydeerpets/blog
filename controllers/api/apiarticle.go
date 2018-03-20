package api

import (
	"blog/controllers"
	"blog/models"
	//"fmt"
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

func (self *ApiarticleController) GetRankList(){
	list := make(map[string]interface{}, 0)

	recommendfilters := make([]interface{}, 0)
	recommendfilters = append(recommendfilters, "is_recommend__exact", 2)
	recommendList, _ := models.ArticleGetList(1, 4, recommendfilters...)

	hotfilters := make([]interface{}, 0)
	hotfilters = append(hotfilters, "is_hot__exact", 2)
	hotList, _ := models.ArticleGetList(1, 4, hotfilters...)

	list["recommendlist"] = recommendList
	list["hotlist"] = hotList

	self.QajaxList("成功", 0, 0, list)
}