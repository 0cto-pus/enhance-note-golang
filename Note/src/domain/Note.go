package domain

import "time"

type Note struct {
	ID        uint64    `json:"id" gorm:"PrimaryKey"`
	UserID    uint64    `json:"user_id" gorm:"index;not null"` // Sadece referans olarak tutulur
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}