package logic

import (
	"GameMail/internal/dao"
	"GameMail/internal/model"
	"GameMail/utils"
	"encoding/json"
	"gorm.io/gorm"
)

type MailOperator struct {
	dao *dao.Dao
}

func NewMailOperator(db *gorm.DB) *MailOperator {
	return &MailOperator{dao: dao.NewDao(db)}
}

func (m *MailOperator) SendEMail(mail *model.Mail) error {
	// 加入邮件
	err := m.dao.AddMail(mail)
	if err != nil {
		utils.Error.Printf("SendEMail call AddMail err=%+v\n", err)
		return err
	}
	// 缓存邮件
	err = dao.GM.SetGlobal(mail)
	if err != nil {
		utils.Error.Printf("SendEMail call SetGlobal err=%+v\n", err)
		return err
	}
	//filter筛选 加入对应用户的邮箱
	var filter *model.FilterRuler
	err = json.Unmarshal([]byte(mail.Filter), &filter)
	if err != nil {
		return err
	}
	switch filter.Field {

	case model.AllUser:
		//发送全体用户
	case model.SpecificUser:
		//发送指定用户
	case model.LevelUser:
		//发送等级用户
	case model.HasGoodUser:
		//发送拥有道具的用户
	}
	return nil

}
