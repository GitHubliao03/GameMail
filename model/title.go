package model

type Title struct {
	TitleID     int    `gorm:"column:title_id" json:"title_id" form:"title_id"`
	TitleString string `gorm:"column:title_string" json:"title_string" form:"title_string"`
}
