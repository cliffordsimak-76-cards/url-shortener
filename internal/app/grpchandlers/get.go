package grpchandlers

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/cliffordsimak-76-cards/url-shortener/internal/proto"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

func (s *GrpcServer) Get(
	ctx context.Context,
	in *pb.GetRequest,
) (*pb.GetResponse, error) {
	url, err := s.repository.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, fmt.Errorf("error get")
		}
		return nil, fmt.Errorf("error get")
	}

	if url.Deleted {
		return nil, nil
	}

	return &pb.GetResponse{
		Url: url.Original,
	}, nil
}
