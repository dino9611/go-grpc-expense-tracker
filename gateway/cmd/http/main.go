package main

import (
	"grpc-finance-app/gateway/cmd/grpcclient"
	"grpc-finance-app/gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	authServiceClient := grpcclient.AuthService()

	h := handler.New(authServiceClient)

	router.POST("register", h.Register)

	router.Run(":8080")
}
