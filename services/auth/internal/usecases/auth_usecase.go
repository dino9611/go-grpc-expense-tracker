package usecases

import (
	"context"
	"fmt"
	authpb "grpc-finance-app/proto"
	"grpc-finance-app/services/auth/internal/dto/req"
	"grpc-finance-app/services/auth/internal/models"
	"grpc-finance-app/services/auth/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthUseCase interface {
	Create(ctx context.Context, authDto *req.AuthReqDto) (*authpb.User, error)
}

type authUseCase struct {
	authRepo repositories.IAuthRepo
}

func NewAuthUseCase(authRepo repositories.IAuthRepo) IAuthUseCase {
	return &authUseCase{
		authRepo: authRepo,
	}
}

func (au *authUseCase) Create(ctx context.Context, authDto *req.AuthReqDto) (*authpb.User, error) {
	hashpass, errbcrypt := bcrypt.GenerateFromPassword([]byte(authDto.Password), 10)

	if errbcrypt != nil {
		return nil, fmt.Errorf("bcript error %v", errbcrypt)
	}
	userData := &models.User{
		Username: authDto.Username,
		Password: string(hashpass),
		Email:    authDto.Email,
	}

	result, err := au.authRepo.AddUser(ctx, userData)

	if err != nil {
		return nil, fmt.Errorf("errror %w", err)
	}

	return result.ToPb(), nil
}
