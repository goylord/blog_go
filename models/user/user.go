package model

type User struct {
	Username string
	Password string
	Id string
	LoginId string
	CreateTime string
	UpdateTime string
}

func (u User) TableName() string {
	return "users"
}