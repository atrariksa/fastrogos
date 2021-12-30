package models

import "time"

type User struct {
	ID        uint
	Username  string `gorm:"index:idx_username,unique"`
	Email     string `gorm:"index:idx_email,unique"`
	Password  string
	Role      string    `gorm:"index:idx_role"`
	UpdatedAt time.Time `gorm:"index:idx_updated_at"`
	CreatedAt time.Time `gorm:"index:idx_created_at"`
}
