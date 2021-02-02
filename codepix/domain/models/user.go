package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID    string `json:"id" valid:"notnull"`
	Name  string `json:"Name" valid:"notnull"`
	Email string `json:"Email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		return err
	}
	return nil
}
func NewUser(name string, email string) (*User, error) {
	user := User{
		ID:    uuid.NewV4().String(),
		Name:  name,
		Email: email,
	}

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
