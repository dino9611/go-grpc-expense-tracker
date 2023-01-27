package grpcclient

import (
	"fmt"
	authpb "grpc-finance-app/proto"
	"log"

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
