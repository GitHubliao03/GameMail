package logic

import (
	"GameMail/utils"
	"gorm.io/gorm"
)

// GetReadMail 获取已读的邮件
func (m *MailOperator) GetReadMail(uid,page,pageSiz int) ([]int, error){
	mailIDs, _, err := m.dao.GetReadMail(uid, page, pageSiz)
	if err != nil {
		utils.Error.Printf("GetReadMail err=%+v\n",err)
		return nil, err
	}
	return mailIDs, nil
}

// GetUnReadMail 获取未读的邮件
func (m *MailOperator)GetUnReadMail(uid,page,pageSize int)([]int,error) {
	mailIDs, _, err := m.dao.GetUnReadMail(uid,page,pageSize)
	if err != nil {
		utils.Error.Printf("GetUnReadMail err=%+v\n",err)
		return nil, err
	}
	return mailIDs,nil
}

// 获取全部的邮件内容
func (m *MailOperator) ReadAllUnReadMail(uid int) {
	// 查询用户对应的未读email id
	var err error
	for err != nil || err != gorm.ErrRecordNotFound {
		err = m.dao.GetUnReadMail()
	}
}