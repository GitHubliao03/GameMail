package model

import (
	"net"
	"time"
)

type User struct {
	ID    int
	Level int
	Name  string

	LastActiveTime time.Time
}

type UserCache struct {
	// 关联的User对象
	User *User

	//客户端与服务器端之间面向流的网络连接
	Conn net.Conn
}
