package models

import "time"

type User struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	Password    string    `gorm:"size:100;not null" json:"password"`
	DOB         time.Time `gorm:"date;" json:"dob"`
	Address     string    `gorm:"size:500;" json:"address"`
	Description string    `gorm:"size:500;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
