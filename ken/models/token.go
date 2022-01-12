package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	ID        uint
	Hash      string    `gorm:"index:idx_hash,unique"`
	Username  string    `gorm:"index:idx_username,unique"`
	IssuedAt  time.Time `gorm:"index:idx_issued_at,unique"`
	ExpiresAt time.Time `gorm:"index:idx_expires_at,unique"`
	UpdatedAt time.Time `gorm:"index:idx_updated_at,unique"`
	CreatedAt time.Time
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
