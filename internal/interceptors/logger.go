package interceptors

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func UnaryServerLogger(l *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()

		l.Info(
			"unary GRPC request started",
			zap.String("method", info.FullMethod),
		)

		m, err := handler(ctx, req)
		if err != nil {
			l.Info(
				"unary gRPC failed with error",
				zap.String("err", err.Error()),
			)
		}

		l.Info(
			"unary GRPC response completed",
			zap.Duration("duration", time.Since(start)),
		)

		return m, err
	}
}

func StreamServerLogger(l *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		l.Info(
			"stream GRPC request started",
			zap.String("method", info.FullMethod),
		)

		err := handler(srv, ss)
		if err != nil {
			l.Info(
				"stream gRPC failed with error",
				zap.String("err", err.Error()),
			)
		}

		l.Info(
			"stream GRPC response completed",
			zap.Duration("duration", time.Since(start)),
		)

		return err
	}
}
