package controllers

import (
	"strings"
	"blog/models"
	"strconv"
	"blog/libs"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	ToolsController
}
//首页
func (self *MainController) Index() {
	self.Data["pageAction"] = "index"
	self.Data["pageTitle"] = "首页"
	self.Layout = "fpublic/layout.html"
	self.TplName = "default/index.html"
}

//文章列表
func (self *MainController) Article(){
	tag, _ := self.GetInt("tag")
	self.Data["pageAction"] = "article"
	self.Data["pageTag"] = tag
	self.Data["pageTitle"] = "文章列表"
	self.Layout = "fpublic/layout.html"
	self.TplName = "default/article.html"
}

func (self *MainController) Detail()  {
	self.Data["Detail"] = true
	self.Data["pageTitle"] = "详情页"
	self.Data["pageAction"] = "detail"

	article_id, _ := self.GetInt("id")
	Article, _ := models.ArticleGetById(article_id)
	ArticleContent, _ := models.ArticleContentGetById(article_id)
	row := make(map[string]interface{})
	row["id"] = Article.Id
	row["title"] = Article.Title
	row["summary"] = Article.Summary
	row["thumb"] = Article.Thumb
	row["status"] = Article.Status
	row["p_sort"] = Article.PSort
	row["author_name"] = Article.AuthorName
	row["scan_count"] = Article.ScanCount
	row["comment_count"] = Article.CommentCount
	row["is_recommend"] = Article.IsRecommend
	row["is_top"] = Article.IsTop
	row["is_hot"] = Article.IsHot
	row["create_time"] = beego.Date(time.Unix(Article.CreateTime, 0), "Y-m-d H:i:s")
	row["update_time"] = beego.Date(time.Unix(Article.UpdateTime, 0), "Y-m-d H:i:s")
	//categoryInfo := getCategoryInfo(categoryList, v.CategoryId)
	//row["category_name"] = categoryInfo.name

	//浏览量+1
	models.Regulate("article", "scan_count", 1, "`Id`=?", article_id)

	self.Data["articleContent"] = ArticleContent
	self.Data["data"] = row

	self.Layout = "fpublic/layout.html"
	self.TplName = "default/detail.html"
}

//时间墙
func (self *MainController) Timeline(){
	self.Data["TimeLine"] = true
	self.Data["pageTitle"] = "时光轴"
	self.Data["pageAction"] = "timeline"

	tag, _ := self.GetInt("tag")
	self.Data["pageTag"] = tag
	self.Layout = "fpublic/layout.html"
	self.TplName = "default/timeline.html"
}

//杂七杂八
func (self *MainController) MixedWork(){
	self.TplName = "default/mixed_work.html"
}

//关于
func (self *MainController) About(){
	self.Data["About"] = true
	self.Data["pageTitle"] = "关于"
	self.Data["pageAction"] = "about"
	self.Layout = "fpublic/layout.html"
	self.TplName = "default/about.html"
}

//获取tag
func (self *MainController) GetTagList(){
	tagList := models.TagGetList()
	self.QajaxList("成功", MSG_OK, 0, tagList)
}

//ajax获取文章列表
func (self *MainController) AjaxArticleTable(){

	page, err := self.GetInt("page")
	if err != nil{
		page = 1
	}

	limit, err := self.GetInt("limit")
	if err != nil{
		limit = 30
	}

	tag, err := self.GetInt("tag")
	if err != nil{
		tag = 0
	}

	//获取taglist
	tagList := models.TagGetList()

	//查询条件
	title := strings.TrimSpace(self.GetString("title"))

	filters := make([]interface{}, 0)

	//filters = append(filters, "status", 2)

	if title != ""{
		filters = append(filters, "title__icontains", title)
	}
	if tag != 0{
		filters = append(filters, "tags__FIND_IN_SET", tag)
	}

	result, count := models.ArticleGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result{
		row := make(map[string]interface{})
		row["id"] = v.Id

		tags := strings.Split(v.Tags, ",")
		tagtemplist := make([]map[string]interface{}, 0)
		for _, v := range tagList {
			tagrow := make(map[string]interface{})
			for i := 0; i < len(tags); i++ {
				tag, _ := strconv.Atoi(tags[i])
				if tag == v.Id {
					tagrow["id"] = v.Id
					tagrow["name"] = v.Name
					tagtemplist = append(tagtemplist,tagrow)
				}
			}
		}

		row["tags"] = tagtemplist
		row["title"] = v.Title
		row["summary"] = v.Summary
		if len(v.Thumb)>0 {
			v.Thumb = libs.GetImageSizetype(v.Thumb, "middle")
		}
		row["thumb"] = v.Thumb
		row["p_sort"] = v.PSort
		row["author_name"] = v.AuthorName
		row["scan_count"] = v.ScanCount
		row["comment_count"] = v.CommentCount
		row["is_recommend"] = v.IsRecommend
		row["is_top"] = v.IsTop
		row["is_hot"] = v.IsHot
		row["create_time"] = v.CreateTime
		row["update_time"] = v.UpdateTime
		//categoryInfo := getCategoryInfo(categoryList, v.CategoryId)
		//row["category_name"] = categoryInfo.name
		list[k] = row
	}
	self.QajaxList("成功", MSG_OK, count, list)

}



