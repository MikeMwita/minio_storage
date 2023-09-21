package grpc_handlers

import (
	pb "github.com/Filtronic/Minio/gapi/pb/mutation_gen"
	"github.com/Filtronic/Minio/util"
	"gorm.io/gorm"
)

type grpcServer struct {
	db *gorm.DB
	pb.UnimplementedMutationServiceServer
	config util.Config
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(config util.Config) (*grpcServer, error) {
	server := &grpcServer{
		config: config,
	}
	return server, nil
}
