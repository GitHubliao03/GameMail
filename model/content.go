package model

type Content struct {
	ContentID     int    `gorm:"column:content_id" json:"content_id" form:"content_id"`
	ContentString string `gorm:"column:content_string" json:"content_string" form:"content_string"`
}

func (c *Content) TableName() string {
	return "Content"
}
