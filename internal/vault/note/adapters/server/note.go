package adapters

import (
	"context"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/vault/note/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/note/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type noteHandlers struct {
	usecases usecases.NoteUsecases
	pb.UnimplementedNoteServiceServer
}

func NewPasswordHandlers(vu usecases.NoteUsecases) *noteHandlers {
	return &noteHandlers{
		usecases: vu,
	}
}

func (nh *noteHandlers) AddNote(ctx context.Context, p *pb.AddNoteRequest) (*pb.AddNoteResponse, error) {
	var response pb.AddNoteResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	note := converter.ToDomainNote(p.Note)

	err = nh.usecases.AddNote(ctx, userID, note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "add note error: %s", err.Error())
	}

	return &response, nil
}

func (nh *noteHandlers) GetNote(ctx context.Context, p *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	var response pb.GetNoteResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	note, err := nh.usecases.GetNote(ctx, userID, p.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get note error: %s", err.Error())
	}

	response.Note = converter.FromDomainNote(note)

	return &response, nil
}

func (nh *noteHandlers) ListNote(ctx context.Context, p *pb.ListNoteRequest) (*pb.ListNoteResponse, error) {
	var response pb.ListNoteResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	noteList, err := nh.usecases.ListNote(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list note error: %s", err.Error())
	}

	response.NoteList = converter.FromDomainNotes(noteList)

	return &response, nil
}
func (nh *noteHandlers) DeleteNote(ctx context.Context, p *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	var response pb.DeleteNoteResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = nh.usecases.DeleteNote(ctx, userID, p.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete note error: %s", err.Error())
	}

	return &response, nil
}
