package main

import (
	"grpc-finance-app/gateway/cmd/grpcclient"
	"grpc-finance-app/gateway/internal/handler"
	"grpc-finance-app/gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	authServiceClient := grpcclient.AuthService()

	h := handler.New(authServiceClient)
	router.Use(middlewares.ErrorMiddleware)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	router.Run(":8080")
}
