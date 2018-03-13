package models

import "github.com/astaxie/beego/orm"

type ArticleContent struct {
	Id int
	ArticleId int
	Content string
}

func (a *ArticleContent) TableName() string {
	return TableName("article_content")
}

func ArticleContentAdd(ac *ArticleContent) (int64, error) {
	return orm.NewOrm().Insert(ac)
}

func (ac *ArticleContent) ArticleContentUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(ac, fields...); err != nil{
		return err
	}
	return nil
}

func ArticleContentGetById(id int) (*ArticleContent, error) {
	a := new(ArticleContent)
	if err := orm.NewOrm().QueryTable(TableName("article_content")).Filter("id", id).One(a); err != nil{
		return nil, err
	}
	return a, nil
}

func (a *ArticleContent) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil{
		return err
	}
	return nil
}