package dao

import (
	"gameMail/model"
	"gameMail/utils"
)

func CheckUserID(uid int) (*model.User, error) {
	user := &model.User{}
	err := utils.Db.Where("uid = ?", uid).First(&user).Error
	return user, err
}
