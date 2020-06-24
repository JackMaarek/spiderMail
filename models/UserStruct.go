package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	gorm.Model
	ID uint64         `gorm:"primary_key"`
	Name string       `gorm:"size:255"`
	Password string   `gorm:"size:255"`
	Email string      `gorm:"size:255"`
	OrganismId uint64
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BeforeSave() error {
	var user User
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func ValidateUser(user User, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Name == "" {
			return errors.New("Required Name")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if user.Name == "" {
			return errors.New("Require Name")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func SaveUser() (*User, error) {

	var err error
	var user User
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func FindAllUsers() (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func FindUserByID(uid uint32) (*User, error) {
	var err error
	var user User
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}

func FindUserByEmail(email string) (*User, error)  {
	var err error
	var user User
	err = db.Debug().Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}