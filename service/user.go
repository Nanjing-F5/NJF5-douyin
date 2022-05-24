package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/models"

	"github.com/golang-module/carbon/v2"
)

type WebUserService struct{}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.AdminWebUserLoginVO) uint64 {
	// var user models.User
	// global.Db.Where("name = ? and password = ?", param.Username, param.Password).First(&user)
	user := dao.UserMgr.Login(param.Username, param.Password)

	return user.Id
}

func (u *WebUserService) Register(param models.AdminWebUserLoginVO) (bool, *models.User) {

	user := dao.UserMgr.GetUserByName(param.Username)

	if user.Id > 0 {
		// 用户已存在
		return false, nil
	} else {
		// 创建新用户
		user.Name = param.Username
		user.Password = param.Password
		user.CreateAt = carbon.Now().ToDateTimeString()
		user.UpdateAt = carbon.Now().ToDateTimeString()
		dao.UserMgr.Register(&user)

		user = dao.UserMgr.GetUserByName(param.Username)

		return true, &user
	}

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
