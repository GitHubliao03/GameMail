package router

import (
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()

	r.Run("192.168.31.63:8888")
}