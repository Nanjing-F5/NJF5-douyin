package service

import (
	"douyin/common"
	"douyin/global"
	"douyin/models"
	"fmt"
)

type WebUserService struct{}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.AdminWebUserLoginVO) uint64 {
	var user models.User
	global.Db.Where("name = ? and password = ?", param.Username, param.Password).First(&user)
	fmt.Printf("登录时：user: %v\n", user)
	return user.Id
}

func (u *WebUserService) GetUserByToken(token string) (models.User, error) {
	var user models.User

	uid, err := common.GetUidByToken(token)
	if err != nil {
		return user, err
	}

	global.Db.Where("id = ?", uid).First(&user)

	fmt.Printf("user: %v\n", user)

	return user, nil
}
