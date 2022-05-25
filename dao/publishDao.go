package dao

import (
	"douyin/global"
	"douyin/models"
)

type PublishManager interface {
	Publish(video *models.Video)
}

type pubManager struct {
}

var PubMgr PublishManager = &pubManager{}

func (pm *pubManager) Publish(video *models.Video) {
	global.Db.Create(video)
}
