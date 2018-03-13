package models

import "github.com/astaxie/beego/orm"

type Auth struct {
	Id         int
	AuthName   string
	AuthUrl    string
	UserId     int
	Pid        int
	Sort       int
	Icon       string
	IsShow     int
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Auth) TableName() string {
	return TableName("uc_auth")
}

func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_auth"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

func AuthAdd(auth *Auth)(int64, error)  {
	return orm.NewOrm().Insert(auth)
}

func (a *Auth) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(a, fields...); err != nil{
		return err
	}
	return nil
}

func AuthGetById(id int) (*Auth, error)  {
	a := new(Auth)

	err := orm.NewOrm().QueryTable(TableName("uc_auth")).Filter("id", id).One(a)
	if err != nil{
		return nil, err
	}
	return a, nil
}


