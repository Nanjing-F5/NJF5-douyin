package initialize

import (
	"douyin/common/middleware"
	"douyin/controller"
	"douyin/global"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	// engine.Static("/video", global.Config.Upload.SavePath)
	engine.Static("/static", "./public")

	// 后台管理员前端接口
	web := engine.Group("/douyin")

	{
		// 用户登录API
		// web.POST("/user/login", controller.Login)
		web.POST("/user/login/", controller.Login)
		web.GET("/feed/", controller.Feed)

		web.GET("/user/", controller.UserInfo)
		// 开启JWT认证，以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
