package service

import (
	"book-category-service/src/model/dto"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type BookCategory interface {
	Create(ctx context.Context, data *dto.CreateBookCategoryReq) error
	FindManyByCategory(ctx context.Context, data *dto.FindManyByCategoryReq) (*dto.DataWithPaging[[]*pb.Book], error)
	Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error
}
