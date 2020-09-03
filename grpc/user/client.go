package grpcuser

import (
	"log"

	"google.golang.org/grpc"

	"cashbag-me-mini/config"
	userpb "cashbag-me-mini/proto"
)

func CreateClient() (*grpc.ClientConn, userpb.UserServiceClient) {
	envVars := config.GetEnv()

	clientConn, err := grpc.Dial(envVars.GRPCUri, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	client := userpb.NewUserServiceClient(clientConn)

	return clientConn, client
}
