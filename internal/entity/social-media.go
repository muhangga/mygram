package entity

import "time"

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primary_key;auto_increment;not null"`
	Name           string    `json:"name" gorm:"unique;not null"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at" gorm:"<-:create, default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"<-:update, default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
