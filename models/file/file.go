package model

type File struct {
	Id string
	AssoId string
	FilePath string
	CreateTime string
	UpdateTime string
}
func (f File) TableName() string {
	return "files"
}