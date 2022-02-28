package controller

import (
	"encoding/json"
	"fmt"
	"gameMail/model"
	"net"
)

//将全体邮件发给所有在线用户

func SendAllMailToOnlineUsers(allMail model.AllMail) {

	//遍历服务器端的onlineUsers map[int]*model.UserCache
	//将消息取出转发

	data, err := json.Marshal(allMail)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//for id, up := range userMgr.onlineUsers {
	//
	//	if id == smsMes.UserId {
	//		continue
	//	}
	//	sp.SendMesToEachOnlineUser(data, up.Conn)
	//}

}

//根据Conn转发数据

func SendMailToEachOnlineUser(data []byte, conn net.Conn) {

	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("转发消息失败，err=", err)
	}
}
