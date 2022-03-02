package model

import (
	"time"
)

type User struct {
	Id             int       `gorm:"column:id;type:int(11)" json:"id"`
	Level          int       `gorm:"column:level" json:"level"`                       // 游戏等级
	Name           string    `gorm:"column:name" json:"name"`                         // 游戏昵称
	LastActiveTime time.Time `gorm:"column:last_active_time" json:"last_active_time"` // 上次活跃时间
	Status         int       `gorm:"column:status" json:"status"`                     // 0上线，1下线
	Identity       int       `gorm:"column:identity" json:"identity"`                 // 0普通玩家，1白名单，2黑名单
	IsDel          int       `gorm:"column:is_del" json:"is_del"`                     // 0未删除，1已删除
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`           // 创建时间
}

func (m *User) TableName() string {
	return "user"
}

type UserCache struct {
	*User
	NextTime time.Time	//下次刷新的时间
	Num int	//刷新缓存的时间
	OnlineTime time.Time	//上线的时间
}

