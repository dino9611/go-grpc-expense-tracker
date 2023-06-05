package models

import (
	"grpc-finance-app/services/auth/internal/dto/res"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"

	"gorm.io/gorm"
)

type User struct {
	ID       int64
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
