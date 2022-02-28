package router

import (
	"gameMail/controller"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()

	r.POST("/get_mail", controller.GetMail)

	r.POST("/logout", controller.Logout)

	r.Run("192.168.31.63:8888")
}
