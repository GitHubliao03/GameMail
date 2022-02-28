package controller

import (
	"gameMail/dao"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetMail(c *gin.Context) {
	//获取用户ID及上次活跃时间

	//返回所有未读未失效邮件
	c.JSON(http.StatusOK, gin.H{
		"allMail":    allMails,
		"activeMail": activeMails,
		"uniqueMail": uniqueMail,
	})
}

func Logout(c *gin.Context) {
	//获取用户ID
	uid := strconv.Atoi(c.PostForm("id"))

	user, err := dao.CheckUserID(uid)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"error": "该用户不存在",
		})
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "服务器内部出错",
		})
	}

}
