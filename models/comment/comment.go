package model

type Comment struct {

}
func (c Comment) TableName() string {
	return "comments"
}