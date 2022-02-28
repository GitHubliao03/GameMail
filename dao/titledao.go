package dao

import (
	"fmt"
	"gameMail/model"
	"gameMail/utils"
	"github.com/jinzhu/gorm"
)

func NewTitle(titleString string) (titleID int) {
	title := model.Title{
		TitleString: titleString,
	}

	utils.Db.Create(&title)

	return title.TitleID
}

//GetTitle 根据标题ID获取标题文本
func GetTitle(titleID int) (string, error) {
	title := &model.Title{}
	err := utils.Db.Where("title_id = ?", titleID).Find(&title).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("没有该标题")
		return "", err
	}
	return title.TitleString, nil
}
