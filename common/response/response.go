package response

import (
	"douyin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	// Data    interface{} `json:"data"`
}

// PageResult 分页结果返回
type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// 用户的登录成功
type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// 返回用户信息
type UserResponse struct {
	Response
	User models.User `json:"user"`
}

// FeedResponse
type FeedResponse struct {
	Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func Success(message string, userId uint64, token string, c *gin.Context) {
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: message},
		UserId:   userId,
		Token:    token,
	})
}

// 返回用户信息
func UserInfo(user models.User, c *gin.Context) {
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: message},
	})
}

// SuccessPage 请求成功返回分页结果
// func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
// 	page := &PageResult{Total: rows, List: data}
// c.JSON(http.StatusOK, Response{200, message, page})
// }
