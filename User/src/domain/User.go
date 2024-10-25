package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}