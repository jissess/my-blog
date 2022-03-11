package models

import (
	"github.com/jinzhu/gorm"
	"my-blog/syserror"
)

type PraiseLog struct {
	Model
	UserId uint
	Key    string
	Table  string
	Flag   bool
}

type TempPraise struct {
	Praise int
}

func UpdatePraise(table string, columnKey string, key string, userId uint) (praise int, err error) {
	d := db.Begin()
	defer func() {
		if err != nil {
			d.Rollback()
		} else {
			d.Commit()
		}
	}()
	var praiseLog PraiseLog
	err = d.Where("`table` = ? and `key` = ? and user_id = ?", table, key, userId).Take(&praiseLog).Error
	if err == gorm.ErrRecordNotFound {
		praiseLog = PraiseLog{
			UserId: userId,
			Key:    key,
			Table:  table,
			Flag:   false,
		}
	} else if err != nil {
		return 0, err
	}
	if praiseLog.Flag {
		return 0, syserror.HasPraiseError{}
	}
	praiseLog.Flag = true
	if err = d.Save(&praiseLog).Error; err != nil {
		return 0, err
	}
	var tp TempPraise
	err = d.Table(table).Where(columnKey+" = ?", key).Select("praise").Scan(&tp).Error
	if err != nil {
		return 0, err
	}
	praise = tp.Praise + 1
	err = d.Table(table).Where(columnKey+" = ?", key).Update("praise", praise).Error
	if err != nil {
		return 0, err
	}
	return praise, nil
}
