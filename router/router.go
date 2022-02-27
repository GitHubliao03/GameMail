package router

import (
	"gameMail/controller"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()

	r.POST("/login", controller.UserLogin)

	r.Run("192.168.31.63:8888")
}
