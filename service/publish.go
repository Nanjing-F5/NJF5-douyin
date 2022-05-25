package service

import (
	"douyin/dao"
	"douyin/models"
)

type PublishService struct{}

// Publish 将用户上传的视频保存到mysql中
func (ps *PublishService) Publish(video *models.Video) {

	dao.PubMgr.Publish(video)
}
