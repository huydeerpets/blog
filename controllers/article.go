package controllers

import (
	"strings"
	"blog/models"
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
)

var (
	AUDIT_STATUS = [5]string{"<span class='layui-badge layui-bg-black'>废弃</span>",
		"<span class='layui-badge layui-bg-orange'>草稿</span>",
		"<span class='layui-badge layui-bg-green'>审核通过</span>"}
	AUDIT_STATUS_TEXT = [5]string{"废弃",
		"草稿",
		"审核通过"}
)

type ArticleController struct {
	BaseController
}

func (self *ArticleController) List(){
	self.Data["pageTitle"] = "文章列表"
	//self.Data["ApiCss"] = true
	self.Data["auditStatus"] = AUDIT_STATUS_TEXT
	self.display()
}

func (self *ArticleController) Table(){
	//列表
	page, err := self.GetInt("page")
	if err != nil{
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil{
		limit = 30
	}

	title := strings.TrimSpace(self.GetString("title"))
	status, _ := self.GetInt("status", -1)

	//获取category
	//categoryList := categoryLists()

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	if status != -1{
		filters = append(filters, "status", status)
	}
	if title != ""{
		filters = append(filters, "title__icontains", title)
	}
	result, count := models.ArticleGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result{
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["summary"] = v.Summary
		row["thumb"] = v.Thumb
		row["status"] = v.Status
		row["status_text"] = AUDIT_STATUS[v.Status]
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
	//fmt.Printf("------------%#v",list)
	self.QajaxList("成功", MSG_OK, count, list)
}

//查看详情
func (self *ArticleController) Detail(){
	self.Data["ApiCss"] = true
	article_id, _ := self.GetInt("id")
	self.Data["pageTitle"] = "查看文章"
	Article, _ := models.ArticleGetById(article_id)
	ArticleContent, _ := models.ArticleContentGetById(article_id)
	row := make(map[string]interface{})
	row["id"] = Article.Id
	row["title"] = Article.Title
	row["summary"] = Article.Summary
	row["thumb"] = Article.Thumb
	row["status"] = Article.Status
	row["status_text"] = AUDIT_STATUS[Article.Status]
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

	self.Data["articleContent"] = ArticleContent
	self.Data["data"] = row
	self.display()
}

//审核
func (self *ArticleController) Audit(){
	
}

func (self *ArticleController) AjaxChangeStatus()  {
	article_id, _ := self.GetInt("id")
	Article, _ := models.ArticleGetById(article_id)

	if Article.Status == 0 || Article.Status == 1{
		Article.Status = 2
	}else{
		Article.Status = 0
	}

	if err := Article.Update(); err != nil{
		self.QajaxMsg(err.Error(), MSG_ERR)
	}
	self.QajaxMsg("", MSG_OK)
}


//新增
func (self *ArticleController) Add(){
	self.Data["pageTitle"] = "添加文章"
	tagList := models.TagGetList()
	self.Data["tagList"] = tagList
	//categoryList := categoryLists()
	//self.Data["categoryList"] = categoryList
	//articleContent := articleContentList()
	//self.Data["articleContent"] = articleContent
	self.display()
}

//修改
func (self *ArticleController) Edit(){
	article_id, _ := self.GetInt("id")
	self.Data["pageTitle"] = "修改文章"
	tagList := models.TagGetList()
	self.Data["tagList"] = tagList
	Article, _ := models.ArticleGetById(article_id)
	ArticleContent, _ := models.ArticleContentGetById(article_id)

	tags := strings.Split(Article.Tags, ",")
	list := make([]map[string]interface{}, len(tagList))
	tagsInitStr := ""
	for k, v := range tagList {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(tags); i++ {
			tag, _ := strconv.Atoi(tags[i])
			if tag == v.Id {
				row["checked"] = 1
				tagsInitStr = tagsInitStr + strconv.Itoa(v.Id) + ","
			}
		}
		row["id"] = v.Id
		row["name"] = v.Name
		list[k] = row
	}
	self.Data["tagsInitStr"] = strings.TrimRight(tagsInitStr,",")
	self.Data["tagList"] = list
	self.Data["articleContent"] = ArticleContent
	self.Data["data"] = Article
	self.display()
}

//根据不同的field修改文章状态
func (self *ArticleController) AjaxUpdate(){
	article_id, _ := self.GetInt("id")
	field_one := strings.TrimSpace(self.GetString("field_one"))
	status, _ := self.GetInt("status")

	_, err := models.FieldUpdateById(article_id, status, field_one)
	if err != nil{
		self.QajaxMsg(err.Error(), MSG_ERR)
	}
	self.QajaxMsg("", MSG_OK)
}

func (self *ArticleController) AjaxSave(){
	article_id, _ := self.GetInt("id")
	//新增
	if article_id == 0{
		Article := new(models.Article)
		Article.Tags = strings.TrimSpace(self.GetString("tags"))
		Article.Title = strings.TrimSpace(self.GetString("title"))
		Article.Summary = strings.TrimSpace(self.GetString("summary"))
		if strings.TrimSpace(self.GetString("thumb")) == ""{
			Article.Thumb = ""
		}else{
			Article.Thumb = beego.AppConfig.String("site.url") + "/" + beego.AppConfig.String("uploadpath") + strings.TrimSpace(self.GetString("thumb"))
		}

		Article.Status, _ = self.GetInt("status")
		Article.AdminId = self.userId
		Article.AuthorName = self.userName
		Article.CreateTime = time.Now().Unix()
		Article.UpdateTime = time.Now().Unix()

		ArticleController := new(models.ArticleContent)

		o := orm.NewOrm()
		err := o.Begin()

		var articleId int64
		articleId, err = o.Insert(Article)
		ArticleController.ArticleId = int(articleId)
		ArticleController.Content = strings.TrimSpace(self.GetString("content"))
		_, err = o.Insert(ArticleController)

		if err != nil{
			err = o.Rollback()
			self.QajaxMsg(err.Error(), MSG_ERR)
		}else{
			err = o.Commit()
			self.QajaxMsg("", MSG_OK)
		}

	}else{
		//修改
		Article, _ := models.ArticleGetById(article_id)
		Article.Tags = strings.TrimSpace(self.GetString("tags"))
		Article.Title = strings.TrimSpace(self.GetString("title"))
		Article.Summary = strings.TrimSpace(self.GetString("summary"))

		if(Article.Thumb != strings.TrimSpace(self.GetString("thumb"))){
			Article.Thumb = beego.AppConfig.String("site.url") + "/" + beego.AppConfig.String("uploadpath") + strings.TrimSpace(self.GetString("thumb"))
		}else{
			Article.Thumb = strings.TrimSpace(self.GetString("thumb"))
		}
		Article.Status, _ = self.GetInt("status")
		Article.UpdateTime = time.Now().Unix()

		ArticleController, err := models.ArticleContentGetById(article_id)

		orm.NewOrm().Begin()

		err = Article.Update()
		ArticleController.Content = strings.TrimSpace(self.GetString("content"))
		err = ArticleController.Update()

		if err != nil{
			err = orm.NewOrm().Rollback()
			self.QajaxMsg(err.Error(), MSG_ERR)
		}else{
			err = orm.NewOrm().Commit()
			self.QajaxMsg("", MSG_OK)
		}
	}

}