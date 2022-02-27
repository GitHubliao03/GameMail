package controller

import (
	"fmt"
	"gameMail/dao"
)

func CreateAllMail(titleString string, contentString string, propID int) error {

	titleID := dao.NewTitle(titleString)
	contentID := dao.NewContent(contentString)

	allMail, err := dao.NewAllMail(titleID, contentID, propID)
	if err != nil {
		fmt.Println("内部服务器错误")
		return err
	}

	//将全体邮件发送给所有在线用户
}
