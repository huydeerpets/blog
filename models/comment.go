package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id int
	ArticleId int
	UserId int
	ReplyTo int
	Nickname string
	Content string
	Status int
	LikeCount int
	CreateTime int64
	UpdateTime int64
}

func (c *Comment) TableName() string {
	return TableName("comment")
}

func  GetCommentByArticleId(page, pageSize int, filters ...interface{}) ([]*Comment, int64)  {
	offset := (page - 1) * pageSize
	list := make([]*Comment, 0)
	query := orm.NewOrm().QueryTable(TableName("comment"))
	if len(filters) > 0{
		l := len(filters)
		for k := 0; k < l; k += 2{
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

