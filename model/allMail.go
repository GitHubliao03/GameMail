package model

import "time"

type AllMail struct {
	MailID     int       `gorm:"column:mail_id" json:"mail_id" form:"mail_id"`
	TitleID    int       `gorm:"column:title_id" json:"title_id" form:"title_id"`
	ContentID  int       `gorm:"column:content_id" json:"content_id" form:"content_id"`
	PropID     int       `gorm:"column:prop_id" json:"prop_id" form:"prop_id"`
	SendDate   time.Time `gorm:"column:send_date" json:"send_date" form:"send_date"`
	ExpiryDate time.Time `gorm:"column:expiry_date" json:"expiry_date" form:"expiry_date"`
}

func (am *AllMail) TableName() string {
	return "AllMails"
}
