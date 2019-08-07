package models

// 用户表
// 说明
type User struct {
	// 用户名称
	Username string `json:"username" build:"post,put,patch"`
	// 密码
	PassWord string `json:"password" bson:"pwd"`
	// 登陆权限
	RightOfLogin bool `json:"right_of_login" build:"post,put,patch interface"`
}
