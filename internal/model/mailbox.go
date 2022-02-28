package model

import (
	"time"
)

type Mailbox struct {
	Uid        int       `gorm:"column:uid" json:"uid"`                 // 用户id
	Eid        int       `gorm:"column:eid" json:"eid"`                 // 邮件id
	Status     int       `gorm:"column:statusL" json:"status"`          // 0未读，1已读，2已读未领取，3已读已领取
	IsDel      int       `gorm:"column:is_del" json:"is_del"`           // 0未删除，1已删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
}

func (m *Mailbox) TableName() string {
	return "mailbox"
}
