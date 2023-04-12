package handler

import (
	"context"
	"fmt"
	authpb "grpc-finance-app/proto"
	"grpc-finance-app/services/auth/internal/dto/req"
	"log"
)

func (h *Handler) Register(ctx context.Context, in *authpb.UserRegisterReq) (*authpb.User, error) {
	authdto := &req.AuthReqDto{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}
	result, err := h.authUseCase.Create(ctx, authdto)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (h *Handler) Login(ctx context.Context, in *authpb.UserLoginReq) (*authpb.User, error) {
	authdto := &req.AuthLoginReqDto{
		Username: in.Username,
		Password: in.Password,
	}

	fmt.Println(authdto)
	result, err := h.authUseCase.Get(ctx, authdto)
	if err != nil {
		log.Println("tes erro")
		return nil, err
	}

	return result, nil

}
