package dao

import (
	"douyin/global"
	"douyin/models"
)

type UserManager interface {
	Register(user *models.User)
	Login(name, password string) models.User
	GetInfoById(uid uint64) models.User
	GetUserByName(name string) models.User
}

type manager struct {
}

var UserMgr UserManager = &manager{}

func (mgr *manager) Login(name, password string) models.User {
	var user models.User

	global.Db.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

func (mgr *manager) Register(user *models.User) {
	global.Db.Create(user)
}

func (mgr *manager) GetInfoById(uid uint64) models.User {
	var user models.User

	global.Db.Where("id = ?", uid).First(&user)
	return user
}

func (mgr *manager) GetUserByName(name string) models.User {
	var user models.User

	global.Db.Where("name = ?", name).First(&user)
	return user
}
