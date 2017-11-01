package models

type User struct {
	Id          int    `id`
	Name        string `用户名`
	PassWord    string `密码`
	MobilePhone string `手机号`
}
