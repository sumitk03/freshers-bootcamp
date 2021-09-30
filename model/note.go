package model

import "time"

type Note struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	NoteTitle     string    `gorm:"type:varchar(10);unique" json:"title"`
	NoteDetail    string    `gorm:"type:varchar(100)" json:"detail"`
	NoteCreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	NoteUpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

type UpdateNote struct {
	NoteDetail string `json:"detail"`
}
