package grpchandlers

import (
	pb "github.com/cliffordsimak-76-cards/url-shortener/internal/proto"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

// GrpcServer.
type GrpcServer struct {
	pb.UnimplementedShortenerServer
	repository repository.Repository
}

// NewGrpcServer.
func NewGrpcServer(
	repository repository.Repository,
) *GrpcServer {
	return &GrpcServer{
		repository: repository,
	}
}
