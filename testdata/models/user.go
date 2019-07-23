package models

type User struct {
	Username string `json:"username"`
	PassWord string `json:"password" bson:"pwd"`
}
