package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id int
	Name string
}

func (t *Tag) TableName() string {
	return TableName("tag")
}

func TagGetList() ([]*Tag)  {
	list := make([]*Tag, 0)
	query := orm.NewOrm().QueryTable(TableName("tag"))
	query.All(&list)
	return list
}


func GetTagInfo(tl []*Tag, tagId int) (tagInfo Tag){
	for _, v := range tl{
		if v.Id == tagId{
			tagInfo = *v
		}
	}
	return
}





