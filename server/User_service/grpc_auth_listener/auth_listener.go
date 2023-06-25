package grpc_auth_listener

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/auth_proto"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/internal/tokenutil"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AuthServer struct {
	auth_proto.UnimplementedAuthServiceServer
}

func (a *AuthServer) IsTokenValid(c context.Context, req *auth_proto.TokenRequest) (*wrapperspb.BoolValue, error) {

	t := strings.Split(req.Token, " ")
	secret_key := os.Getenv("SECRET_KEY")
	if len(t) == 2 {
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret_key)
		if authorized {
			return &wrapperspb.BoolValue{Value: true}, nil
		}

		if err != nil {
			log.Println(err)
			return &wrapperspb.BoolValue{Value: false}, err
		}

	}
	return &wrapperspb.BoolValue{Value: false}, nil

}

func GRPCListen() {
	gRPCPort := 9002
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
