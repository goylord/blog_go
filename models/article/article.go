package model

type Articles struct {
	Id string
	Title string
	Content string
	CreateTime string
	UpdateTime string
	ClassId string
	CreatorId string
	PreviewContent string
}
func (a Articles) TableName() string {
	return "articles"
}