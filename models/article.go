package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int
	Tags string
	Title string
	Summary string
	Thumb string
	Status int
	PSort int
	AdminId int
	AuthorName string
	ScanCount int
	CommentCount int
	IsRecommend int
	IsTop int
	IsHot int
	CreateTime int64
	UpdateTime int64
}

func (a *Article) TableName() string {
	return TableName("article")
}

func ArticleAdd(a *Article) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func (a *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil{
		return err
	}
	return nil
}

func FieldUpdateById(id int, isyes int, fieldone string) (num int64, err error){
	sql := "UPDATE qq_article set " + fieldone + " =? where id = ?"
	res, err := orm.NewOrm().Raw(sql, isyes, id).Exec()
	num = 0
	if err == nil{
		num, _ = res.RowsAffected()
	}
	return num, err
}

func ArticleGetById(id int) (*Article, error) {
	a := new(Article)
	if err := orm.NewOrm().QueryTable(TableName("article")).Filter("id", id).One(a); err != nil{
		return nil, err
	}
	return a, nil
}

func ArticleGetList(page, pageSize int, filters ...interface{}) ([]*Article, int64)  {
	offset := (page - 1) * pageSize
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(TableName("article"))
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

