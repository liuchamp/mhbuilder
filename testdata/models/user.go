package models

// 用户表
// 说明
type User struct {
	// 用户名称
	Username string `json:"username" build:"post,put,patch"`
	// 密码
	PassWord string `json:"password" bson:"pwd" scope:"198"`
	// 登陆权限
	RightOfLogin bool `json:"right_of_login" build:"post,put,patch,filter" scope:"129"`
	Finder       bool `json:"finder" build:"post,put,patch" scope:"199"`
	Code         bool `json:"finder" build:"put,patch,filter" scope:"199"`
}

type Finder struct {
	Uid string `json:"uid" build:"post,patch,filter" scope:"199"`
	Mid string `json:"mid" build:"post,patch,filter" scope:"25"`
}
