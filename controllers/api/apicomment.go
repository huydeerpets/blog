package api

import (
	"blog/controllers"
	"blog/models"
	"strings"
	"github.com/astaxie/beego/orm"
	"time"
)

type ApicommentController struct {
	controllers.ToolsController
}

func (self *ApicommentController) GetComment(){
	articleId, err := self.GetInt("aid")
	if err != nil{
		self.QajaxMsg(err.Error(), -1)
	}

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
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["content"] = v.Content
		row["create_time"] = v.CreateTime
		row["like_count"] = v.LikeCount
		row["reply_to"] = v.ReplyTo
		row["nickname"] = v.Nickname
		itemfilters := make([]interface{}, 0)
		itemfilters = append(itemfilters, "article_id", articleId)
		itemfilters = append(itemfilters, "reply_to", v.Id)
		itemCommentRes, _ := models.GetCommentByArticleId(1, 100, itemfilters...)
		item := make ([]map[string]interface{}, len(itemCommentRes))
		for kk, vv := range itemCommentRes{
			itemrow := make(map[string]interface{})
			itemrow["id"] = vv.Id
			itemrow["content"] = vv.Content
			itemrow["create_time"] = vv.CreateTime
			itemrow["like_count"] = v.LikeCount
			itemrow["nickname"] = v.Nickname
			item[kk] = itemrow
		}
		row["item"] = item
		list[k] = row
	}
	self.QajaxList("成功", 0, count, list)
}

func (self *ApicommentController) PostComment(){
	article_id, err := self.GetInt("aid")
	replay_to, err := self.GetInt("reply_to", 0)
	user_id, _ := self.GetInt("uid")
	content := strings.TrimSpace(self.GetString("content"))
	if err != nil{
		self.QajaxMsg(err.Error(), -1)
	}
	if len(content) <= 0 {
		self.QajaxMsg("请填写内容", -1)
	}
	mc := new(models.Comment)
	mc.ArticleId = article_id
	mc.Content = content
	mc.UserId = user_id
	mc.ReplyTo = replay_to
	mc.CreateTime = time.Now().Unix()
	_, err = orm.NewOrm().Insert(mc)
	//if replay_to != 0 && err == nil {
	//	one := models.Comment{Id: replay_to}
	//	one.LikeCount = one.LikeCount+1
	//	orm.NewOrm().Update(&one, "like_count")
	//}
	if err != nil{
		self.QajaxList(err.Error(), -1, 0, "")
	}
	self.QajaxList("成功", 0, 0, mc)
}

//评论点赞

