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
		PlayUrl:       "http://10.136.240.231:8080/static/bear.mp4",
		CoverUrl:      "http://10.136.240.231:8080/static/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoUser = models.User{
	Id:   1,
	Name: "TestUser",
}
