package models

type Message struct {
	Model
	MessageKey string `gorm:"size:50;unique_index;not null" json:"message_key"`
	Content    string `json:"content"`
	UserId     uint   `json:"user_id"`
	User       User   `json:"user"`
	NoteKey    string `gorm:"size:50" json:"note_key"`
	Praise     int    `gorm:"default:0" json:"praise"`
}

func SaveMessage(message *Message) error {
	return db.Save(message).Error
}

func QueryMessageByNoteKey(noteKey string) (msg []*Message, err error) {
	return msg, db.Where("note_key = ?", noteKey).Preload("User").Order("created_at desc").Find(&msg).Error
}

func QueryMessageCountByNoteKey(noteKey string) (count int64, err error) {
	return count, db.Model(&Message{}).Where("note_key = ?", noteKey).Count(&count).Error
}

func QueryPageMessageByNoteKey(noteKey string, page, pageSize int) (ms []*Message, err error) {
	return ms, db.Where("note_key = ?", noteKey).Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&ms).Error
}
