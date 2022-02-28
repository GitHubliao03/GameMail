package model

import (
	"time"
)

type Mail struct {
	Id         int       `gorm:"column:id" json:"id"`                   // 邮件id
	Content    string    `gorm:"column:content" json:"content"`         // 邮件的内容
	SendTime   time.Time `gorm:"column:send_time" json:"send_time"`     // 邮件发送的时间
	ClearTime  time.Time `gorm:"column:clear_time" json:"clear_time"`   // 被清理的时间
	Filter     string    `gorm:"column:filter" json:"filter"`           // 过滤条件
	IsDel      int       `gorm:"column:is_del" json:"is_del"`           // 0未删除，1已删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建的时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新的时间
}

func (m *Mail) TableName() string {
	return "mail"
}

// MailContent 邮件的内容
type MailContent struct {
	Title string `json:"title"` //标题
	Context string `json:"context"` //正文
	GoodsID string `json:"goods_id"`//领取的物品id
}