package controller

import (
	"fmt"
	"gameMail/model"
)

//UserMgr实例在服务器端有且只有一个
//在许多地方都会用到，因此将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*model.UserCache
}

//onlineUsers初始化
func init() {

	fmt.Println("onlineUsers初始化")
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*model.UserCache, 1024),
	}

}

//添加在线用户

func (um *UserMgr) AddOnlineUser(up *model.UserCache) {
	fmt.Println("添加在线用户")
	um.onlineUsers[up.User.UserID] = up
}

//删除在线用户

func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

//返回当前所有在线用户

func (um *UserMgr) GetAllOnlineUser() map[int]*model.UserCache {
	return um.onlineUsers
}

//根据id返回对应的值

func (um *UserMgr) GetOnlineUserById(userId int) (up *model.UserCache, err error) {

	//从map中取出一个值，带检测方法
	up, ok := um.onlineUsers[userId]
	if !ok {
		//说明该Id用户不在线
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return
}
