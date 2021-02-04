package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User represents the pixKey user
type User struct {
	Base  `valid:"notnull"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
