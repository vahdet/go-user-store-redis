package main

import (
	"fmt"
	"github.com/vahdet/go-user-store/dal"
	"github.com/vahdet/go-user-store/services"
	"net"
	pb "github.com/vahdet/go-user-store/proto"
	"google.golang.org/grpc"
	"github.com/vahdet/go-user-store/grpcserver"
	"google.golang.org/grpc/reflection"
	log "github.com/sirupsen/logrus"
)

const (
	defaultGrpcPort = "5300"
)

func main() {

	// Initialize DataStore
	dal.InitDataStoreClient()
	defer dal.CloseDataStoreClient()

	// Serve GRPC
	lis, err := net.Listen("tcp", defaultGrpcPort)
	if err != nil {
		panic(fmt.Errorf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &grpcserver.UserServer{Service: services.NewUserService(dal.NewUserDal())})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
