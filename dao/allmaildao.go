package dao

import (
	"fmt"
	"gameMail/model"
	"gameMail/utils"
	"time"
)

func NewAllMail(titleID int, contentID int, propID int) (*model.AllMail, error) {
	allMail := &model.AllMail{
		TitleID:    titleID,
		ContentID:  contentID,
		PropID:     propID,
		SendDate:   time.Now(),
		ExpiryDate: time.Now().AddDate(0, 0, 7),
	}

	err := utils.Db.Create(allMail).Error
	if err != nil {
		fmt.Println("new all mail failed:", err)
		return nil, err
	} else {
		fmt.Println("new all mail success:", err)
		return allMail, nil
	}
}
