package domain

import "time"

type Suggestion struct {
	ID        uint64    `json:"id" gorm:"PrimaryKey"`
	UserID    uint64    `json:"user_id" gorm:"index;not null"` // Sadece referans olarak tutulur
	NoteID uint64 `json:"note_id" gorm:"index;not null"`
	Suggestion   string    `json:"suggestion" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
} 