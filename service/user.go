package service

import (
	"douyin/global"
	"douyin/models"
	"fmt"
)

type WebUserService struct{}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.AdminWebUserLoginVO) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	fmt.Printf("user.Id: %v\n", user.Id)
	return user.Id
}
