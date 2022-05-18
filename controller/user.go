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

	// fmt.Printf("param: %v\n", param)

	uid := userService.Login(param)

	// uid>0 说明在数据库中能查找到该用户
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username, uid)
		// userInfo := models.AdminWebUserInfo{
		// 	Uid:   uid,
		// 	Token: token,
		// }
		response.Success("登录成功", uid, token, c)
		fmt.Println("登录成功...")
		return
	}
	response.Failed("User doesn't exist", c)
}

// 获取当前登录用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user, err := userService.GetUserByToken(token)
	if err != nil {
		response.Failed(err.Error(), c)
	}

	// user = models.User{
	// 	Id:            1,
	// 	Name:          "kkite",
	// 	FollowCount:   12,
	// 	FollowerCount: 10000,
	// 	IsFollow:      false,
	// }

	if user.Id > 0 {
		response.UserInfo(user, c)
		return
	}
	response.Failed("User doesn't exist", c)

}
