package dao

import (
	"fmt"
	"gameMail/model"
	"gameMail/utils"
	"github.com/jinzhu/gorm"
)

func NewContent(contentString string) (titleID int) {
	content := model.Content{
		ContentString: contentString,
	}

	utils.Db.Create(&content)

	return content.ContentID
}

//GetContent 根据内容ID获取内容文本
func GetContent(contentID int) (string, error) {
	content := &model.Content{}
	err := utils.Db.Where("content_id = ?", contentID).Find(&content).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("没有该标题")
		return "", err
	}
	return content.ContentString, nil
}
