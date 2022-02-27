package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	//获取用户ID及最大全体邮件ID

	//返回所有未读未失效邮件
	c.JSON(http.StatusOK, gin.H{
		"allMail":    allMails,
		"activeMail": activeMails,
		"uniqueMail": uniqueMail,
	})
}
