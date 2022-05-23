package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/models"
)

type WebUserService struct{}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.AdminWebUserLoginVO) uint64 {
	// var user models.User
	// global.Db.Where("name = ? and password = ?", param.Username, param.Password).First(&user)
	user := dao.UserMgr.Login(param.Username, param.Password)

	return user.Id
}

func (u *WebUserService) GetUserByToken(token string) (models.User, error) {
	var user models.User

	uid, err := common.GetUidByToken(token)
	if err != nil {
		return user, err

	}

	user = dao.UserMgr.GetInfoById(uid)

	return user, nil
}
