package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	// fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Admin), new(Auth), new(Role), new(RoleAuth),new(Article),new(Tag),new(ArticleContent), new(Comment))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

//设置字段值增减
//@param                table           需要操作的数据表
//@param                field           需要对值进行增减的字段
//@param                step            增减的步长，正值为加，负值为减
//@param                condition       查询条件
//@param                conditionArgs   查询条件参数
//@return               err             返回错误
func Regulate(table, field string, step int, condition string, conditionArgs ...interface{}) (err error) {
	table = TableName(table) //表处理
	mark := "+"             //符号
	if step < 0 {           //步长处理
		step = -step
		mark = "-"
	}
	sql := fmt.Sprintf("update %v set %v=%v%v? where %v", table, field, field, mark, condition)

	if len(conditionArgs) > 0 {
		_, err = orm.NewOrm().Raw(sql, step, conditionArgs[0:]).Exec()
	} else {
		_, err = orm.NewOrm().Raw(sql, step).Exec()
	}
	return err
}