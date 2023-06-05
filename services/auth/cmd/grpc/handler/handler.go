package handler

import (
	"grpc-finance-app/services/auth/internal/usecases"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"
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
