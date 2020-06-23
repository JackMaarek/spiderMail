package Models

import "time"

type Token struct {
	ID        uint64 `gorm:"primary_key"`
	Token     string `gorm:"size:255"`
	ExpiresAt time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	Revoked   bool
	User      User
}