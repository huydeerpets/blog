package models

type Comment struct {
	Id int
	ArticleId int
	AdminId int
	ReplyTo int
	NickName string
	Email string
	Content string
	Status int
	CreateTime int
	UpdateTime int
}

func (c *Comment) TableName() string {
	return TableName("comment")
}
