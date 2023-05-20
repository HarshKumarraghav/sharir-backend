package auth

import "github.com/google/uuid"

type User struct {
	Id          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Password    string `json:"password,omitempty" validate:"required,min=6"`
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Profile     string `json:"profile,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	UserName    string `json:"userName,omitempty" validate:"required"`
}
type OutUser struct {
	Id          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Profile     string `json:"profile,omitempty"`
	Name        string `json:"name,omitempty"`
	UserName    string `json:"userName,omitempty"`
}
type InUser struct {
	Email       string `json:"email,omitempty" validate:"required,email"`
	Password    string `json:"password,omitempty" validate:"required,min=6"`
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Profile     string `json:"profile,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	UserName    string `json:"userName,omitempty" validate:"required"`
}

func (in *InUser) ToUser() *User {
	UUID := uuid.New().String()
	return &User{
		Id:          UUID,
		Email:       in.Email,
		Password:    in.Password,
		PhoneNumber: in.PhoneNumber,
		Profile:     in.Profile,
		Name:        in.Name,
		UserName:    in.UserName,
	}
}
func (u *User) ToOutUser() *OutUser {
	return &OutUser{
		Id:          u.Id,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Profile:     u.Profile,
		Name:        u.Name,
		UserName:    u.UserName,
	}
}
func (u *User) ToInUser() *InUser {
	return &InUser{
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
		Profile:     u.Profile,
		Name:        u.Name,
		UserName:    u.UserName,
	}
}
