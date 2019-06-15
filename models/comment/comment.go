package model

type Comment struct {
	Id string
	ArticleId string 
	ParentId  string
	CreatorId string
	Content   string
	CreateTime string
	UpdateTime string
}
func (c Comment) TableName() string {
	return "comments"
}