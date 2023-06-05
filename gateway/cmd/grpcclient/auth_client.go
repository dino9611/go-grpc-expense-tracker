package grpcclient

import (
	"fmt"
	"log"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthService() authpb.AuthsClient {
	port := "20002"

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("could not connect to :", port, err)
	}

	return authpb.NewAuthsClient(conn)
}
