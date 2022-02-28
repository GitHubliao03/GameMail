package controller

import (
	"encoding/json"
	"fmt"
	"gameMail/model"
	"net"
)

//将全体邮件发给所有在线用户

func SendAllMailToOnlineUsers(allMail *model.AllMail) error {

	//遍历服务器端的onlineUsers map[int]*model.UserCache
	//将消息取出转发

	data, err := json.Marshal(allMail)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}

	for _, up := range userMgr.onlineUsers {
		err = SendMailToEachOnlineUser(data, up.Conn)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据Conn转发数据

func SendMailToEachOnlineUser(data []byte, conn net.Conn) error {

	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("转发消息失败，err=", err)
		return err
	}

	return nil
}
