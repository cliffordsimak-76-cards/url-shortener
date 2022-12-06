package grpchandlers

import (
	pb "github.com/cliffordsimak-76-cards/url-shortener/internal/proto"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

type GrpcServer struct {
	pb.UnimplementedShortenerServer
	repository repository.Repository
}

func NewGrpcServer(
	repository repository.Repository,
) *GrpcServer {
	return &GrpcServer{
		repository: repository,
	}
}
