package handler

import (
	"grpc-finance-app/services/auth/internal/usecases"

	authpb "grpc-finance-app/proto"
)

type Handler struct {
	authUseCase usecases.IAuthUseCase
	authpb.UnimplementedAuthsServer
}

func New(authUseCase usecases.IAuthUseCase) *Handler {
	return &Handler{
		authUseCase: authUseCase,
	}
}
