package models

import (
	"errors"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/jinzhu/gorm"
	"time"
)

type Token struct {
	ID        uint64    `gorm:"primary_key"`
	Token     string    `gorm:"size:255"`
	ExpiresAt time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	Revoked   bool      `gorm:"default:FALSE"`
	UserId    uint64
}

func CreateTokenFromUser(user *User) (*Token, error) {
	var token string
	var err error
	token, err = services.CreateToken(user.ID)
	if err != nil {
		return &Token{}, err
	}
	var expireDate time.Duration
	expireDate, err = services.CreateTokenExpireDate()
	if err != nil {
		return &Token{}, err
	}

	var createdToken = Token{
		Token:     token,
		ExpiresAt: time.Now().Add(expireDate).UTC(),
		Revoked:   false,
		UserId:    user.ID,
	}

	err = db.Debug().Create(&createdToken).Error
	if err != nil {
		return &Token{}, err
	}
	return &createdToken, err
}

func FindTokenByUserID(uid uint64) (*Token, error) {
	var err error
	var token Token
	err = db.Debug().Model(&Token{}).Where("user_id = ?", uid).Take(&token).Error
	if err != nil {
		return &Token{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Token{}, errors.New("Token Not Found")
	}
	return &token, nil
}

func FindTokenByToken(token string) (*Token, error) {
	var err error
	var tokenStruct Token
	err = db.Debug().Model(&Token{}).Where("token = ?", token).Take(&tokenStruct).Error
	if err != nil {
		return &Token{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Token{}, errors.New("Token not fount")
	}
	return &tokenStruct, nil
}

func UpdateToken(token string) error {
	var err error
	var oldToken *Token
	var newExpireDate time.Duration
	var newToken string
	newExpireDate, err = services.CreateTokenExpireDate()
	if err != nil {
		return err
	}
	oldToken, err = FindTokenByToken(token)
	if err != nil {
		return err
	}
	newToken, err = services.CreateToken(oldToken.UserId)
	if err != nil {
		return err
	}
	oldToken.Token = newToken
	oldToken.ExpiresAt = time.Now().Add(newExpireDate)
	err = db.Debug().Save(&oldToken).Error
	if err != nil {
		return err
	}
	return nil
}
