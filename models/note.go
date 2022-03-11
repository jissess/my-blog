package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Note struct {
	Model
	NoteKey string `gorm:"size:50;unique_index;not null"`
	UserId  int
	User    User
	Title   string `gorm:"type:varchar(200)"`
	Summary string `gorm:"type:varchar(800)"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("note_key = ?", key).Take(&note).Error
}

func QueryNoteByKeyAndUserId(key string, userId uint) (note Note, err error) {
	return note, db.Where("note_key = ? and user_id = ?", key, userId).Take(&note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}

func QueryNotesByPage(title string, page, limit int) (note []*Note, err error) {
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("id DESC").Find(&note).Error
}

func QueryNotesCount(title string) (count int, err error) {
	return count, db.Model(&Note{}).Count(&count).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Error
}

func DeleteNote(note *Note) error {
	return db.Delete(note).Error
}

func DeleteNoteByKeyAndUserId(key string, userId uint) error {
	return db.Delete(&Note{}, "note_key = ? and user_id = ?", key, userId).Error
}

func NoteVisitInc(key string) error {
	return db.Model(&Note{}).Where("note_key = ?", key).UpdateColumn("visit", gorm.Expr("visit + ?", 1)).Error
}
