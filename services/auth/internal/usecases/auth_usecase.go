package usecases

import (
	"context"
	"fmt"
	"grpc-finance-app/services/auth/internal/dto/req"
	"grpc-finance-app/services/auth/internal/errs"
	"grpc-finance-app/services/auth/internal/models"
	"grpc-finance-app/services/auth/internal/repositories"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type IAuthUseCase interface {
	Create(ctx context.Context, authDto *req.AuthReqDto) (*authpb.User, error)
	Get(ctx context.Context, authDto *req.AuthLoginReqDto) (*authpb.User, error)
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
		return nil, fmt.Errorf("error from usecase (hash) : %w", errs.ErrorHashBcrypt)
	}
	userData := &models.User{
		Username: authDto.Username,
		Password: string(hashpass),
		Email:    authDto.Email,
	}

	result, err := au.authRepo.AddUser(ctx, userData)

	if err != nil {
		return nil, fmt.Errorf("error from repo : %w", err)
	}

	return result.ToPb(), nil
}

func (au *authUseCase) Get(ctx context.Context, authDto *req.AuthLoginReqDto) (*authpb.User, error) {
	result, err := au.authRepo.GetUserByUsername(ctx, authDto.Username)

	if err != nil {
		return nil, fmt.Errorf("error from repo : %w", err)
	}

	isPassMatch := CheckPasswordHash(authDto.Password, result.Password)

	if !isPassMatch {
		return nil, fmt.Errorf("error from usecase (hash) : %w", errs.ErrorPassWrong)
	}

	return result.ToPb(), nil

}
