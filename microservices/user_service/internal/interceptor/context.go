package interceptor

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TgIDFromContext(ctx context.Context) (int64, error) {
	tgID, ok := ctx.Value(ctxKey{}).(int64)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "tg_id not found in context")
	}
	return tgID, nil
}
