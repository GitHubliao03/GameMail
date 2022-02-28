package dao

import (
	"fmt"
	"gameMail/model"
	"gameMail/utils"
)

func CheckUserID(userID int) (*model.User, error) {
	user := &model.User{}
	err := utils.Db.Where("user_id = ?", userID).First(&user).Error
	return user, err
}

//UpdateUser 将在线用户缓存写入数据库
func UpdateUser(userID int, userCache *model.UserCache) error {
	user := &model.User{}
	err := utils.Db.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return err
	}
	user = userCache.User

	err = utils.Db.Where("user_id = ?", userID).Save(&user).Error
	if err != nil {
		fmt.Println("用户更新失败")
		return err
	}
	return nil
}
