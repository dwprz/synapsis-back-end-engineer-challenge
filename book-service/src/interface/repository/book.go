package repository

import (
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type Book interface {
	Create(ctx context.Context, data *dto.AddBookReq) error
	FindByTitle(ctx context.Context, title string) (*entity.Book, error)
	FindManyByTitle(ctx context.Context, title string, limit, offset int) ([]*entity.Book, error)
	FindManyByFields(ctx context.Context, data *dto.GetBookReq, limit, offset int) ([]*entity.Book, error)
	FindManyByRandom(ctx context.Context, limit, offset int) ([]*entity.Book, error)
	FindManyByIds(ctx context.Context, ids []uint32) ([]*pb.Book, error)
	UpdateById(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error)
	DeleteById(ctx context.Context, bookId int) error
}
