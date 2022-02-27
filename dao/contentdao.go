package dao

import (
	"gameMail/model"
	"gameMail/utils"
)

func NewContent(contentString string) (titleID int) {
	content := model.Content{
		ContentString: contentString,
	}

	utils.Db.Create(&content)

	return content.ContentID
}
