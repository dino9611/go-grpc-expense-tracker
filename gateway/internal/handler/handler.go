package handler

import (
	authpb "grpc-finance-app/proto"
)

type Handler struct {
	authServiceClient authpb.AuthsClient
}

func New(aac authpb.AuthsClient) *Handler {
	return &Handler{
		authServiceClient: aac,
	}
}
