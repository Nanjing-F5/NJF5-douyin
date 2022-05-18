package controller

import (
	"net/http"
	"time"

	"douyin/common/response"
	"douyin/global"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: global.VideoList,
		NextTime:  time.Now().Unix(),
	})
}
