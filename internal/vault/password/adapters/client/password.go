package adapters

import (
	"context"
	"errors"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault"
	"github.com/Chystik/pass-man/internal/vault/password/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/password/entities"

	"google.golang.org/grpc"
)

type passwordAPIClient struct {
	conn     *grpc.ClientConn
	password pb.PasswordServiceClient
	vault.PasswordAPIClient
}

func NewPasswordAPIClient(conn *grpc.ClientConn, password pb.PasswordServiceClient) *passwordAPIClient {
	return &passwordAPIClient{
		conn:     conn,
		password: password,
	}
}

func (vc *passwordAPIClient) AddPassword(ctx context.Context, p entities.Password) error {
	req := &pb.AddPasswordRequest{
		Password: converter.FromDomainPassword(p),
	}

	res, err := vc.password.AddPassword(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}
func (vc *passwordAPIClient) GetPassword(ctx context.Context, meta string) (entities.Password, error) {
	p := entities.Password{}

	req := &pb.GetPasswordRequest{
		Meta: meta,
	}

	res, err := vc.password.GetPassword(ctx, req)
	if err != nil {
		return p, err
	}

	if res.Error != nil {
		return p, errors.New(res.Error.String())
	}

	return converter.ToDomainPassword(res.Password), nil
}
func (vc *passwordAPIClient) ListPassword(ctx context.Context) ([]entities.Password, error) {
	p := []entities.Password{}
	req := &pb.ListPasswordRequest{}

	res, err := vc.password.ListPassword(ctx, req)
	if err != nil {
		return p, err
	}

	if res.Error != nil {
		return p, errors.New(res.Error.String())
	}

	return converter.ToDomainPasswords(res.PasswordList), nil
}

func (vc *passwordAPIClient) DeletePassword(ctx context.Context, meta string) error {
	req := &pb.DeletePasswordRequest{
		Meta: meta,
	}

	res, err := vc.password.DeletePassword(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}
