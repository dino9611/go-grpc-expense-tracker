package main

import (
	"fmt"
	"grpc-finance-app/services/auth/cmd/grpc/handler"
	"grpc-finance-app/services/auth/internal/config"
	"grpc-finance-app/services/auth/internal/db"
	"grpc-finance-app/services/auth/internal/repositories"
	"grpc-finance-app/services/auth/internal/usecases"
	"log"
	"net"
	"os"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.InitConfig()
	db := db.ConnectPgDb(cfg)
	authRepo := repositories.NewAuthRepo(db)
	authUseCase := usecases.NewAuthUseCase(authRepo)
	handler := handler.New(authUseCase)

	s := grpc.NewServer()
	authpb.RegisterAuthsServer(s, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Application.Port))

	if err != nil {
		os.Exit(1)
	}
	log.Println("Starting RPC server at", cfg.Application.Port)
	if err != nil {
		log.Fatalf("could not listen to %v: %v", cfg.Application.Port, err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %w", err)
		os.Exit(1)
	}
}
