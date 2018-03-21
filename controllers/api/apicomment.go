package api

import (
	"blog/controllers"
	"blog/models"
)

type ApicommentController struct {
	controllers.ToolsController
}

func (self *ApicommentController) GetComment(){
	articleId, _ := self.GetInt("aid")

	page, err := self.GetInt("page")
	if err != nil{
		page = 1
	}

	limit, err := self.GetInt("limit")
	if err != nil{
		limit = 30
	}

	filters := make([]interface{}, 0)

	filters = append(filters, "article_id", articleId)
	filters = append(filters, "reply_to", 0)

	parentCommentRes, count := models.GetCommentByArticleId(page, limit, filters...)
	list := make([]map[string]interface{}, len(parentCommentRes))

	for k, v := range parentCommentRes{

	}


}
