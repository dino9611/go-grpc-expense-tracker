package models

import (
	authpb "grpc-finance-app/proto"
	"grpc-finance-app/services/auth/internal/dto/res"

	"gorm.io/gorm"
)

type User struct {
	ID       int
	Email    string `gorm:"unique"`
	Password string
	Username string `gorm:"unique"`
	*gorm.Model
}

func (u *User) ToDto() *res.AuthResDto {
	userdto := &res.AuthResDto{
		Id:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	return userdto
}

func (u *User) ToPb() *authpb.User {
	userpb := &authpb.User{
		Id:       int64(u.ID),
		Username: u.Username,
		Email:    u.Email,
	}
	return userpb
}
