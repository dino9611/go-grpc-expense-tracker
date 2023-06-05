package handler

import (
	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"
)

type Handler struct {
	authServiceClient authpb.AuthsClient
}

func New(aac authpb.AuthsClient) *Handler {
	return &Handler{
		authServiceClient: aac,
	}
}
