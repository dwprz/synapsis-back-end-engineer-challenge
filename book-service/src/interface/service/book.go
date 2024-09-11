package service

import (
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type Book interface {
	Add(ctx context.Context, data *dto.AddBookReq) error
	FindMany(ctx context.Context, data *dto.GetBookReq) (*dto.DataWithPaging[[]*entity.Book], error)
	FindManyByIds(ctx context.Context, data *dto.FindManyByIdsReq) ([]*pb.Book, error)
	FindManyPopularBook(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error)
	Delete(ctx context.Context, bookId int) error
}
