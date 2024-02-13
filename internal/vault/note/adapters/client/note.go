package adapters

import (
	"context"
	"errors"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/note/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/note/entities"

	"google.golang.org/grpc"
)

type NoteAPIClient interface {
	AddNote(ctx context.Context, note entities.Note) error
	GetNote(ctx context.Context, meta string) (entities.Note, error)
	ListNote(ctx context.Context) ([]entities.Note, error)
	DeleteNote(ctx context.Context, meta string) error
}

type noteAPIClient struct {
	conn *grpc.ClientConn
	note pb.NoteServiceClient
	NoteAPIClient
}

func NewNoteAPIClient(conn *grpc.ClientConn, note pb.NoteServiceClient) *noteAPIClient {
	return &noteAPIClient{
		conn: conn,
		note: note,
	}
}

func (nc *noteAPIClient) AddNote(ctx context.Context, note entities.Note) error {
	req := &pb.AddNoteRequest{
		Note: converter.FromDomainNote(note),
	}

	res, err := nc.note.AddNote(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}

func (nc *noteAPIClient) GetNote(ctx context.Context, meta string) (entities.Note, error) {
	n := entities.Note{}

	req := &pb.GetNoteRequest{
		Meta: meta,
	}

	res, err := nc.note.GetNote(ctx, req)
	if err != nil {
		return n, err
	}

	if res.Error != nil {
		return n, errors.New(res.Error.String())
	}

	return converter.ToDomainNote(res.Note), nil
}

func (nc *noteAPIClient) ListNote(ctx context.Context) ([]entities.Note, error) {
	n := []entities.Note{}

	req := &pb.ListNoteRequest{}

	res, err := nc.note.ListNote(ctx, req)
	if err != nil {
		return n, err
	}

	if res.Error != nil {
		return n, errors.New(res.Error.String())
	}

	return converter.ToDomainNotes(res.NoteList), nil
}

func (nc *noteAPIClient) DeleteNote(ctx context.Context, meta string) error {
	req := &pb.DeleteNoteRequest{
		Meta: meta,
	}

	res, err := nc.note.DeleteNote(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}
