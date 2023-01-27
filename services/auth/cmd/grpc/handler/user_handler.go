package handler

import (
	"context"
	authpb "grpc-finance-app/proto"
	"grpc-finance-app/services/auth/internal/dto/req"
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
