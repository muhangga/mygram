package entity

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment;not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Age       uint8     `json:"age"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create, default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update, default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
