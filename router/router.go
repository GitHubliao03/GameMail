package router

import (
	"gameMail/controller"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()

	//获取所有未读邮件
	r.GET("/mail", controller.GetMail)

	//登陆游戏加入在线用户缓存
	r.POST("/login", controller.Login)

	//退出游戏记录最后活跃时间
	r.POST("/logout", controller.Logout)

	//获取标题、内容
	r.GET("/title", controller.GetTitle)
	r.GET("/content", controller.GetContent)

	r.Run("192.168.31.63:8888")
}
