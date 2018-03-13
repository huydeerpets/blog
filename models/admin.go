package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Admin struct{
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Admin) TableName() string{
	return TableName("uc_admin")
}

func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error){
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("login_name", loginName).One(a)
	if err != nil{
		return nil, err
	}
	return a, nil
}

func AdminGetById(id int) (*Admin, error){
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("id", id).One(a)
	if err != nil{
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{})([]*Admin, int64){
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_admin"))
	if len(filters) > 0 {
		l := len(filters)
		fmt.Printf("-----%#v", filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.Limit(pageSize, offset).All(&list)

	return list, total
}

func (a *Admin) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(a, fields...); err != nil{
		return err
	}
	return nil
}