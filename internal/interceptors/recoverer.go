package interceptors

import (
	"runtime/debug"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerRecoverer(l *zap.Logger) grpc.UnaryServerInterceptor {
	var grpcPanicRecoveryHandler = func(p any) error {
		l.Error(
			"recovered from panic",
			zap.Any("panic", p),
			zap.String("stack", string(debug.Stack())),
		)
		return status.Errorf(codes.Internal, "%s", p)
	}
	return recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler))
}

func StreamServerRecoverer(l *zap.Logger) grpc.StreamServerInterceptor {
	var grpcPanicRecoveryHandler = func(p any) error {
		l.Error(
			"recovered from panic",
			zap.Any("panic", p),
			zap.String("stack", string(debug.Stack())),
		)
		return status.Errorf(codes.Internal, "%s", p)
	}
	return recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler))
}
