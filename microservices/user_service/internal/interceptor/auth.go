package interceptor

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey struct{}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authValues := md["authorization"]
	if len(authValues) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	token := strings.TrimPrefix(authValues[0], "Bearer ")

	tgID, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	ctx = context.WithValue(ctx, ctxKey{}, tgID)

	return handler(ctx, req)
}
