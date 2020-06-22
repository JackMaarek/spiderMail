package Models

import "time"

type Token struct {
	ID uint64 `json:"id"`
	Token string `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User User
}