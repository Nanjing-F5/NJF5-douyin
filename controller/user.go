package controller

import (
	"douyin/common"
	"douyin/common/response"
	"douyin/models"
	"douyin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var userService service.WebUserService

// Login 后台管理前端，用户登录
func Login(c *gin.Context) {
	var param models.AdminWebUserLoginVO

	param.Username = c.Query("username")
	param.Password = c.Query("password")

	uid := userService.Login(param)

	// uid>0 说明在数据库中能查找到该用户
	if uid > 0 {
		token, err := common.GenerateToke(param.Username, uid)
		if err != nil {
			response.Failed("Generate token failed", c)
			return
		}

		response.Success("登录成功", uid, token, c)
		fmt.Println("登录成功...")
		return
	}
	response.Failed("User doesn't exist", c)
}

// 响应当前登录用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user, err := userService.GetUserByToken(token)
	if err != nil {
		response.Failed(err.Error(), c)
		return
	}

	if user.Id > 0 {
		response.UserInfo(user, c)
		return
	}
	response.Failed("User doesn't exist", c)
}

// 用户注册
func Register(c *gin.Context) {
	var param models.AdminWebUserLoginVO

	fmt.Println("111")
	param.Username = c.Query("username")
	param.Password = c.Query("password")

	flag, user := userService.Register(param)

	fmt.Println("flag: ", flag)

	if flag {
		// 返回创建user的response
		token, err := common.GenerateToke(user.Name, user.Id)
		if err != nil {
			response.Failed("Generate token failed", c)
			return
		}
		response.Success("注册成功", user.Id, token, c)
	} else {
		// 返回user已存在的response
		response.Failed("User already exist", c)
	}
}
