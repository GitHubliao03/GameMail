package logic

import (
	"GameMail/internal/dao"
	"GameMail/internal/model"
	"GameMail/utils"
	"encoding/json"
	"gorm.io/gorm"
)

// GetReadMail 获取已读的邮件
func (m *MailOperator) GetReadMail(uid, page, pageSiz int) ([]int, error) {
	mailIDs, _, err := m.dao.GetReadMail(uid, page, pageSiz)
	if err != nil {
		utils.Error.Printf("GetReadMail err=%+v\n", err)
		return nil, err
	}
	return mailIDs, nil
}

// GetUnReadMail 获取未读的邮件
func (m *MailOperator) GetUnReadMail(uid, page, pageSize int) ([]int, error) {
	mailIDs, _, err := m.dao.GetUnReadMail(uid, page, pageSize)
	if err != nil {
		utils.Error.Printf("GetUnReadMail err=%+v\n", err)
		return nil, err
	}
	return mailIDs, nil
}

// ReadAllUnReadMail 获取全部的邮件内容
func (m *MailOperator) ReadAllUnReadMail(uid int) ([]int, error) {
	// 查询用户对应的未读email id
	var err error
	var i int
	emailIDs := make([]int, 0, 20)
	for ; ; i++ {
		eids, _, err := m.dao.GetUnReadMail(uid, i, 20)
		if err != nil {
			break
		}
		emailIDs = append(emailIDs, eids...)
	}
	if err != gorm.ErrRecordNotFound {
		utils.Error.Printf("ReadAllUnReadMail call GetUnReadMail err=%+v\n", err)
		return nil, err
	}
	// 查找一系列的邮件
	mailsContent := make([]string, 0, len(emailIDs))
	mailIDs := make([]int, 0, len(emailIDs))
	for _, eids := range emailIDs {
		global, err := dao.GM.GetGlobal(eids)
		if err != nil {
			mailIDs = append(mailIDs, eids)
			continue
		}
		v, _ := global.(*model.Mail)
		mailsContent = append(mailsContent, v.Content)
	}
	mc, _, err := m.dao.GetMailContentByIDs(mailIDs)
	if err != nil {
		utils.Error.Printf("ReadAllUnReadMail call GetMailByIDs err=%+v\n", err)
		return nil, err
	}
	mailsContent = append(mailsContent, mc...)
	//领取道具
	goodIds := make([]int,0,len(emailIDs))
	for _,v:= range mailsContent {
		var content model.MailContent
		err := json.Unmarshal([]byte(v), &content)
		if err != nil {
			utils.Error.Printf("ReadAllUnReadMail call json.Unmarshal err=%+v\n",err)
			return nil, err
		}
		goodIds = append(goodIds,content.GoodsID)
	}
	// 把一系列邮件更新为已读已领取
	_, err = m.dao.UpdateMailBoxStatusBatch(mailIDs,uid,3)
	if err != nil {
		utils.Error.Printf("ReadAllUnReadMail call UpdateMailBoxStatusBatch err=%+v\n",err)
		return nil, err
	}
	return goodIds, err
}
