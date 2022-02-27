package dao

import (
	"gameMail/model"
	"gameMail/utils"
)

func NewTitle(titleString string) (titleID int) {
	title := model.Title{
		TitleString: titleString,
	}

	utils.Db.Create(&title)

	return title.TitleID
}
