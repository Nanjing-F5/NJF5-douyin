package models

// 数据库，用户数据映射模型
type User struct {
	Id       uint64 `gorm:"primayKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Phone    string `gorm:"phone"`
	CreateAt string `gorm:"create_at"`
	UpdateAt string `gorm:"update_at"`
}

// 后台管理用户登录值对象ValueObject，用来传递参数
type AdminWebUserLoginVO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 后台管理前端，用户信息传输模型
type AdminWebUserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}
