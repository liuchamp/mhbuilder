package models

type User struct {
	Username string `json:"username" build:"post,put,patch"`
	PassWord string `json:"password" bson:"pwd"`
}
