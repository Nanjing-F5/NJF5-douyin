package controller

import (
	"douyin/common"
	"douyin/common/response"
	"douyin/models"
	"douyin/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user service.WebUserService

// Login 后台管理前端，用户登录
func Login(c *gin.Context) {
	var param models.AdminWebUserLoginVO

	param.Username = c.Query("username")
	param.Password = c.Query("password")

	// fmt.Printf("param: %v\n", param)

	uid := user.Login(param)

	// uid>0 说明在数据库中能查找到该用户
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		// userInfo := models.AdminWebUserInfo{
		// 	Uid:   uid,
		// 	Token: token,
		// }
		response.Success("登录成功", uid, token, c)
		fmt.Println("登录成功...")
		return
	}
	response.Failed("用户名或密码错误", c)
}

// 只是测试
func UserInfo(c *gin.Context) {
	// token := c.Query("token")

	user := models.User{
		Id:            1,
		Name:          "kkite",
		FollowCount:   12,
		FollowerCount: 10000,
		IsFollow:      false,
	}

	c.JSON(http.StatusOK, response.UserResponse{
		Response: response.Response{StatusCode: 0},
		User:     user,
	})
}
