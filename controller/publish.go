package controller

import (
	"douyin/common/response"
	"douyin/global"
	"douyin/models"
	"douyin/service"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

var publishService service.PublishService

func Publish(c *gin.Context) {
	token := c.PostForm("token")

	user, err := userService.GetUserByToken(token)
	if err != nil {
		response.Failed("User doesn't exist", c)
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		response.Failed(err.Error(), c)
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)

	err = c.SaveUploadedFile(data, saveFile)
	if err != nil {
		response.Failed(err.Error(), c)
		return
	}

	// 将视频保存到数据库中
	video := &models.Video{
		AuthorId: user.Id,
		PlayUrl:  global.Config.Upload.VideoUrl + finalName,
		CoverUrl: global.Config.Upload.CoverUrl,
		CreateAt: carbon.Now().ToDateTimeString(),
	}

	publishService.Publish(video)

	response.PublishSuccess(finalName, c)

}
