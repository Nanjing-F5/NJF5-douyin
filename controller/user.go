package controller

import (
	"douyin/common"
	"douyin/common/response"
	"douyin/models"
	"douyin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var user service.WebUserService

// Login 后台管理前端，用户登录
func Login(c *gin.Context) {
	var param models.AdminWebUserLoginVO
	username, _ := c.GetPostForm("username")
	fmt.Printf("username: %v\n", username)
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	fmt.Printf("param: %v\n", param)
	// 生成token
	uid := user.Login(param)

	fmt.Printf("uid: %v\n", uid)

	// uid>0 说明在数据库中能查找到该用户
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		userInfo := models.AdminWebUserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}
