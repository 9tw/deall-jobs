package repository

import (
	"deall/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname" form:"fullname" validate:"required,min=3,max=20"`
	Email    string `gorm:"unique" json:"email" form:"email"  validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,containsany=1234567890" `
	Role     uint   `json:"role" validate:"numeric"`
	Token    string `json:"token" validate:"multibyte"`
}

type Fullname struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required,min=3,max=20"`
}

type Email struct {
	Email string `gorm:"unique" json:"email" form:"email"  validate:"required,email"`
}

type Password struct {
	Password string `json:"password" form:"password" validate:"required,min=8,containsany=1234567890" `
}

func FromDomain(du domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Fullname: du.Fullname,
		Email:    du.Email,
		Password: du.Password,
		Role:     du.Role,
		Token:    du.Token,
	}
}

func ToDomain(u User) domain.UserCore {
	return domain.UserCore{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
		Token:    u.Token,
	}
}

func ToDomainArray(u []User) []domain.UserCore {
	var res []domain.UserCore
	for _, val := range u {
		res = append(res, domain.UserCore{ID: val.ID, Fullname: val.Fullname, Email: val.Email, Role: val.Role})
	}

	return res
}
