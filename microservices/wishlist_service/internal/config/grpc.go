package config

import (
	"fmt"
	"log"
	"os"
)

type GRPCConfig interface {
	UserServiceAddress() string
	ServiceAddress() string
	HTTPAddress() string
}

type grpcConfig struct {
	userServiceAddress string
	serviceAddress     string
	httpAddress        string
}

func NewGRPCConfig() GRPCConfig {
	userServiceHost := os.Getenv("USER_SERVICE_GRPC_HOST")
	if userServiceHost == "" {
		userServiceHost = "localhost"
	}

	userServicePort := os.Getenv("USER_SERVICE_GRPC_PORT")
	if userServicePort == "" {
		userServicePort = "3030"
	}

	if userServiceHost == "" || userServicePort == "" {
		log.Fatal("USER_SERVICE_GRPC_HOST and USER_SERVICE_GRPC_PORT must be configured")
	}

	serviceHost := os.Getenv("GRPC_HOST")
	if serviceHost == "" {
		serviceHost = "0.0.0.0"
	}

	servicePort := os.Getenv("GRPC_PORT")
	if servicePort == "" {
		servicePort = "8080"
	}

	httpHost := os.Getenv("HTTP_HOST")
	if httpHost == "" {
		httpHost = "0.0.0.0"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	return &grpcConfig{
		userServiceAddress: fmt.Sprintf("%s:%s", userServiceHost, userServicePort),
		serviceAddress:     fmt.Sprintf("%s:%s", serviceHost, servicePort),
		httpAddress:        fmt.Sprintf("%s:%s", httpHost, httpPort),
	}
}

func (g *grpcConfig) UserServiceAddress() string {
	return g.userServiceAddress
}

func (g *grpcConfig) ServiceAddress() string {
	return g.serviceAddress
}

func (g *grpcConfig) HTTPAddress() string {
	return g.httpAddress
}
