package dao

import (
	"GameMail/internal/model"

	"time"
)

// AddMail 新增一封邮件
func (dao *Dao) AddMail(mail *model.Mail) error {
	return dao.db.Create(&mail).Error
}

// UpdateMail 更新邮件
func (dao *Dao) UpdateMail(id int, fields map[string]interface{}) error {
	return dao.db.Where("id=?", id).Updates(&fields).Error
}

// FindMailByID 查找邮件通过id
func (dao *Dao) FindMailByID(id int) (*model.Mail, error) {
	var mail *model.Mail
	return mail, dao.db.Where("id=?", id).First(&mail).Error
}

// DeleteMail 删除邮件
func (dao *Dao) DeleteMail(id int) error {
	return dao.db.Where("id=?", id).Update("is_del", 1).Error
}

// DeleteMailBatch 批量删除邮件
func (dao *Dao) DeleteMailBatch(clearTime time.Time) (int64, error) {
	db := dao.db.Where("clear_time>?", clearTime).Update("is_del", 1)
	return db.RowsAffected, db.Error
}

// GetMailContentByIDs 查看一系列指定的邮件的内容
func (dao *Dao)GetMailContentByIDs(ids []int)([]string,int64,error) {
	var mailsContent []string
	db := dao.db.Select("content").
		Where("id IN ? and is_del = 0 and send_time > ?",ids,time.Now()).Find(&mailsContent)
	return mailsContent,db.RowsAffected,db.Error
}