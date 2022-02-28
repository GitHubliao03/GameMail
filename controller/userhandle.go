package controller

import (
	"fmt"
	"gameMail/dao"
	"gameMail/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetMail(c *gin.Context) {
	//获取用户ID及上次活跃时间

	//返回所有未读未失效邮件
	//c.JSON(http.StatusOK, gin.H{
	//	"allMail":    allMails,
	//	"activeMail": activeMails,
	//	"uniqueMail": uniqueMail,
	//})
}

//Login 用户登陆，将其加入在线用户缓存
func Login(c *gin.Context) {
	//获取用户ID
	userID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数出错",
		})
		return
	}

	user, err := dao.CheckUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "服务器内部错误",
		})
	}

	userCache := &model.UserCache{
		User: user,
		//Conn:链接谁来获取及维护
	}

	userMgr.AddOnlineUser(userCache)
}

//Logout 用户退出登陆，保存最后活跃时间，并将缓存写入数据库，将其从在线用户列表中删除
func Logout(c *gin.Context) {
	//获取用户ID
	userID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数出错",
		})
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == userID {
			up.User.LastActiveTime = time.Now()
		}
	}
	//每个用户退出登陆后将缓存写入数据库并将其从在线用户缓存中删掉
	userCache, err := userMgr.GetOnlineUserById(userID)
	if err != nil {
		fmt.Println("该账号已经退出登陆")
		c.JSON(http.StatusOK, gin.H{
			"error": "服务器内部错误",
		})
		return
	}
	err = dao.UpdateUser(userID, userCache)
	if err != nil {
		fmt.Println("用户更新失败")
		c.JSON(http.StatusOK, gin.H{
			"error": "服务器内部错误",
		})
		return
	}
	userMgr.DelOnlineUser(userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "user logout success",
	})
}

//UpdateAllOnlineUser 定时落库，将所有在线用户缓存写入数据库
func UpdateAllOnlineUser() error {

	for userId, _ := range userMgr.onlineUsers {
		userCache, err := userMgr.GetOnlineUserById(userId)
		if err != nil {
			fmt.Println("服务器内部错误")
			return err
		}
		err = dao.UpdateUser(userId, userCache)
		if err != nil {
			fmt.Println("用户更新失败")
			return err
		}
	}
	return nil
}

func GetTitle(c *gin.Context) {

	titleID, err := strconv.Atoi(c.PostForm("title_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数出错",
		})
		return
	}

	titleString, err := dao.GetTitle(titleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "没有该标题",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title_string": titleString,
		})
	}
}

func GetContent(c *gin.Context) {

	contentID, err := strconv.Atoi(c.PostForm("content_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数出错",
		})
		return
	}

	contentString, err := dao.GetContent(contentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "没有该内容",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title_string": contentString,
		})
	}
}
