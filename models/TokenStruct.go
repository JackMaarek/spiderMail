package models

import (
	"github.com/JackMaarek/spiderMail/services"
	"time"
)

type Token struct {
	ID        uint64 `gorm:"primary_key"`
	Token     string `gorm:"size:255"`
	ExpiresAt time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	Revoked   bool
	UserId    uint64
}

func CreateUserToken(user *User) (*Token, error) {
	var token string
	var err error
	token, err = services.CreateToken(user.ID)
	var createdToken = Token{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 1),
		Revoked:   false,
		UserId:      user.ID,
	}

	err = db.Debug().Create(&createdToken).Error
	if err != nil {
		return &Token{}, err
	}
	return nil, err
}