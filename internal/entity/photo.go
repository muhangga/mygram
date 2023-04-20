package entity

import (
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment;not null"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create, default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update, default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
