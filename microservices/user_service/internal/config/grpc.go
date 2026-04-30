package config

import (
	"fmt"
	"os"
)

type GRPCConfig interface {
	ServiceAddress() string
}

type grpcConfig struct {
	serviceAddress string
}

func NewGRPCConfig() GRPCConfig {
	serviceHost := os.Getenv("GRPC_HOST")
	if serviceHost == "" {
		serviceHost = "0.0.0.0"
	}

	servicePort := os.Getenv("GRPC_PORT")
	if servicePort == "" {
		servicePort = "8050"
	}

	return &grpcConfig{
		serviceAddress: fmt.Sprintf("%s:%s", serviceHost, servicePort),
	}
}

func (g *grpcConfig) ServiceAddress() string {
	return g.serviceAddress
}
