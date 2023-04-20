package entity

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment;not null"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create, default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update, default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
