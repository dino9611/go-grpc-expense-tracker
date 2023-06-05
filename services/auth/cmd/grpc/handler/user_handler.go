package handler

import (
	"context"
	"errors"
	"grpc-finance-app/services/auth/internal/dto/req"
	"grpc-finance-app/services/auth/internal/errs"
	"grpc-finance-app/services/auth/internal/utils"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (h *Handler) Register(ctx context.Context, in *authpb.UserRegisterReq) (*authpb.User, error) {
	authdto := &req.AuthReqDto{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}
	result, err := h.authUseCase.Create(ctx, authdto)
	if err != nil {
		if errors.Is(err, errs.ErrorUsernameExist) {
			err = status.Error(codes.AlreadyExists, "user already exist")
			return nil, err
		} else if errors.Is(err, errs.ErrorUsernameExist) {
			err = status.Error(codes.Internal, "db error")
			return nil, err
		}
		err = status.Error(codes.Internal, "internal error")
		return nil, err
	}

	return result, nil

}

func (h *Handler) Login(ctx context.Context, in *authpb.UserLoginReq) (*authpb.User, error) {
	authdto := &req.AuthLoginReqDto{
		Username: in.Username,
		Password: in.Password,
	}

	result, err := h.authUseCase.Get(ctx, authdto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, errs.ErrorPassWrong) {
			err = status.Error(codes.NotFound, "user not found")
			return nil, err
		}
		err = status.Error(codes.Internal, "internal error")
		return nil, err
	}

	err = utils.CreateToken(result)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, errs.ErrorPassWrong) {
			err = status.Error(codes.NotFound, "user not found")
			return nil, err
		}
		err = status.Error(codes.Internal, "internal error")
		return nil, err
	}

	return result, nil

}
