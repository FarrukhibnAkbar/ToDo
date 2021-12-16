package grpc_client

import (
	"github.com/FarrukhibnAkbar/ToDo/config"
)

type IGrpcClient interface{}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
