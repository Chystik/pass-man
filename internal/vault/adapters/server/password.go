package adapters

import (
	"context"
	"errors"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/vault/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type vaultHandlers struct {
	usecases usecases.PasswordUsecases
	pb.UnimplementedPasswordServiceServer
}

func NewPasswordHandlers(vu usecases.PasswordUsecases) *vaultHandlers {
	return &vaultHandlers{
		usecases: vu,
	}
}

func (vh *vaultHandlers) AddPassword(ctx context.Context, p *pb.AddPasswordRequest) (*pb.AddPasswordResponse, error) {
	var response pb.AddPasswordResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	password := converter.ToDomainPassword(p.Password)

	err = vh.usecases.AddPassword(ctx, userID, password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "add password error: %s", err.Error())
	}

	return &response, nil
}

func (vh *vaultHandlers) GetPassword(ctx context.Context, p *pb.GetPasswordRequest) (*pb.GetPasswordResponse, error) {
	var response pb.GetPasswordResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	password, err := vh.usecases.GetPassword(ctx, userID, p.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get password error: %s", err.Error())
	}

	response.Password = converter.FromDomainPassword(password)

	return &response, nil
}

func (vh *vaultHandlers) ListPassword(ctx context.Context, p *pb.ListPasswordRequest) (*pb.ListPasswordResponse, error) {
	var response pb.ListPasswordResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	passwordList, err := vh.usecases.ListPassword(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list password error: %s", err.Error())
	}

	response.PasswordList = converter.FromDomainPasswords(passwordList)

	return &response, nil
}

func (vh *vaultHandlers) DeletePassword(ctx context.Context, p *pb.DeletePasswordRequest) (*pb.DeletePasswordResponse, error) {
	var response pb.DeletePasswordResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = vh.usecases.DeletePassword(ctx, userID, p.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete password error: %s", err.Error())
	}

	return &response, nil
}

func getUserFromContext(ctx context.Context) (string, error) {
	claims, ok := ctx.Value(entities.ClaimsKeyName).(*entities.AuthClaims)
	if !ok {
		return "", errors.New("bad auth claims")
	}

	return claims.Login, nil
}
