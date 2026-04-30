package config

import (
	"errors"
	"fmt"
	"os"
)

type GRPCConfig interface {
	ServiceAddress() string
}

type grpcConfig struct {
	serviceAddress string
}

func NewGRPCConfig() (GRPCConfig, error) {
	serviceHost := os.Getenv("GRPC_HOST")
	if serviceHost == "" {
		return nil, errors.New("GRPC_HOST environment variable not set")
	}

	servicePort := os.Getenv("GRPC_PORT")
	if servicePort == "" {
		return nil, errors.New("GRPC_PORT environment variable not set")
	}

	return &grpcConfig{
		serviceAddress: fmt.Sprintf("%s:%s", serviceHost, servicePort),
	}, nil
}

func (g *grpcConfig) ServiceAddress() string {
	return g.serviceAddress
}
