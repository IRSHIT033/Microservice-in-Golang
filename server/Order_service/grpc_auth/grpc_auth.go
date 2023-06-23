package grpc_auth

import (
	"context"
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/auth_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CheckTokenValidityViaGRPC(token string) bool {
	conn, err := grpc.Dial("localhost:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return false
	}
	defer conn.Close()

	client := auth_proto.NewAuthServiceClient(conn)

	response, err := client.IsTokenValid(context.Background(), &auth_proto.TokenRequest{Token: token})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return false
	}

	return response.Value
}
