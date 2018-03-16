package controllers

import (
	"strings"
	"blog/models"
	"fmt"
	"strconv"
)

type MainController struct {
	ToolsController
}

//首页
func (self *MainController) Index() {
	self.TplName = "default/index.html"
}

//文章列表
func (self *MainController) Article(){

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

//获取tag
func (self *MainController) GetTagList(){
	tagList := models.TagGetList()
	self.ajaxList("成功", MSG_OK, 0, tagList)
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

	//获取taglist
	//tagList := models.TagGetList()

	//查询条件
	title := strings.TrimSpace(self.GetString("title"))
	tag := strings.TrimSpace(self.GetString("tag"))

	filters := make([]interface{}, 0)

	//filters = append(filters, "status", 2)

	if title != ""{
		filters = append(filters, "title__icontains", title)
	}
	if tag != ""{
		filters = append(filters, "tags__icontains", tag)
	}
	result, count := models.ArticleGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result{
		row := make(map[string]interface{})
		row["id"] = v.Id

		tagsIdArr := strings.Split(v.Tags, ",")
		fmt.Printf("a-------------------%#v", tagsIdArr)
		for _, vv := range tagsIdArr{
			vv= int(vv)
			fmt.Printf("b--------------%#v", strconv.Atoi(vv))
			//vv, _ = strconv.Atoi(vv)
			//models.GetTagInfo(tagList, vv)
		}

		row["tags"] = v.Tags
		row["title"] = v.Title
		row["summary"] = v.Summary
		row["thumb"] = v.Thumb
		row["p_sort"] = v.PSort
		row["author_name"] = v.AuthorName
		row["san_count"] = v.ScanCount
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
	self.ajaxList("成功", MSG_OK, count, list)

}



