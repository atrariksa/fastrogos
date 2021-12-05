package models

import "time"

type User struct {
	ID        uint
	UserID    string `gorm:"index:idx_user_id,unique"`
	Username  string `gorm:"index:idx_username,unique"`
	Email     string `gorm:"index:idx_email,unique"`
	Address   string
	Password  string
	UpdatedAt time.Time `gorm:"index:idx_updated_at"`
	CreatedAt time.Time `gorm:"index:idx_created_at"`
}
