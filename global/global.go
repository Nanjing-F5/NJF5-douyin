package global

import (
	"douyin/common/config"
	"douyin/models"

	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)

var VideoList = []models.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.137.1:8080/static/bear.mp4",
		CoverUrl:      "http://192.168.137.1:8080/static/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []models.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = models.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
