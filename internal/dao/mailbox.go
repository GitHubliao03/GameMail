package dao

import "GameMail/internal/model"

// AddMailBoxOne 邮箱新增一个
func (dao *Dao) AddMailBoxOne(mailBox *model.Mailbox) error {
	return dao.db.Create(&mailBox).Error
}

// AddMailBoxBatch 批量新增邮件
func (dao *Dao) AddMailBoxBatch(mailBoxes []*model.Mailbox) (int64, error) {
	db := dao.db.Table("mailbox").Create(&mailBoxes)
	return db.RowsAffected, db.Error
}

// UpdateMailBoxStatus 更新邮件
func (dao *Dao) UpdateMailBoxStatus(uid int, eid int, fields map[int]*model.Mailbox) error {
	return dao.db.Table("mailbox").Where("id=? and eid = ?", uid, eid).
		Updates(&fields).Error
}

// UpdateMailBoxStatusBatch 批量更新邮件
func (dao *Dao) UpdateMailBoxStatusBatch( eids []int,uid,status int) (int64, error) {
	db := dao.db.Table("mailbox").Where("uid = ? AND eid IN ?", uid, eids).Update("status", status)
	return db.RowsAffected, db.Error
}

// DeleteMailByID 删除邮件
func (dao *Dao) DeleteMailByID(uid, eid int) error {
	return dao.db.Table("mailbox").Where("uid = ? AND eid = ?",uid,eid).
		Update("is_del",1).Error
}

// GetUnReadMail 查看未读邮件
func (dao *Dao)GetUnReadMail(uid,page,pageSize int,) ([]int, int64, error) {
	var eids []int
	db := dao.db.Table("mailbox").Select("uid = ? and status = 0",uid).
		Offset(page-1).Limit(pageSize).Find(&eids)
	return eids,db.RowsAffected,db.Error
}

// GetReadMail 查看已读邮件
func (dao *Dao)GetReadMail(uid,page,pageSize int) ([]int, int64, error) {
	var eids []int
	db := dao.db.Table("mailbox").Select("uid = ? and status >= 1 ",uid).
		Offset(page-1).Limit(pageSize).Find(&eids)
	return eids,db.RowsAffected,db.Error
}
