package models

type User struct {
	Username     string `json:"username" build:"post,put,patch"`
	PassWord     string `json:"password" bson:"pwd"`
	RightOfLogin bool   `json:"right_of_login" build:"post,put,patch interface"`
}
