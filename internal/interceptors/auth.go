package interceptors

import (
	"context"
	"errors"

	"github.com/Chystik/pass-man/internal/user/entities"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type WrappedServerStream struct {
	grpc.ServerStream
	WrappedContext context.Context
}

func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}
	return &WrappedServerStream{ServerStream: stream, WrappedContext: stream.Context()}
}

func UnaryServerAuth(jwtKey []byte) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		m := info.FullMethod

		// Client token requesting
		if m == "/pb.UserService/Login" || m == "/pb.UserService/SignUp" {
			return handler(ctx, req)
		}

		var tokenStr string
		var ctxToken context.Context

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			values := md.Get("token")
			if len(values) > 0 {
				tokenStr = values[0]
			}
		}
		if len(tokenStr) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		claims := &entities.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return nil, status.Error(codes.Unauthenticated, "invalid token signature")
			}
			return nil, status.Error(codes.Unauthenticated, "bad token")
		}

		if token.Valid {

			// Access context values in handlers like this
			// props, _ := r.Context().Value(entities.ClaimsKeyName).(*jwt.MapClaims)
			ctxToken = context.WithValue(ctx, entities.ClaimsKeyName, claims)
		} else {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		return handler(ctxToken, req)
	}

}

func StreamServerAuth(jwtKey []byte) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		m := info.FullMethod

		// Client token requesting
		if m == "/pb.UserService/Login" || m == "/pb.UserService/SignUp" {
			return handler(srv, ss)
		}

		var tokenStr string
		var ctxToken context.Context
		ctx := ss.Context()

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			values := md.Get("token")
			if len(values) > 0 {
				tokenStr = values[0]
			}
		}
		if len(tokenStr) == 0 {
			return status.Error(codes.Unauthenticated, "missing token")
		}

		claims := &entities.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return status.Error(codes.Unauthenticated, "invalid token signature")
			}
			return status.Error(codes.Unauthenticated, "bad token")
		}

		if token.Valid {
			ctxToken = context.WithValue(ctx, entities.ClaimsKeyName, claims)
		} else {
			return status.Error(codes.Unauthenticated, "invalid token")
		}

		ssWrapped := &WrappedServerStream{
			ServerStream:   ss,
			WrappedContext: ctxToken,
		}

		return handler(srv, ssWrapped)
	}
}
