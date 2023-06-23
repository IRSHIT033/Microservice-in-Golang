package grpc_auth_listener

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/auth_proto"
	"google.golang.org/grpc"
)

type AuthServer struct {
	auth_proto.UnimplementedAuthServiceServer
}

func (a *AuthServer) SendProduct(c context.Context, req *auth_proto.TokenRequest) (bool, error) {
	return false, nil
}

func GRPCListen() {
	gRPCPort := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()
	auth_proto.RegisterAuthServiceServer(s, &AuthServer{})
	log.Printf("grpc server started on port %d", gRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC:%v", err)
	}
}
