package service

import (
	"book-category-service/src/api/grpc/client"
	"book-category-service/src/common/errors"
	"book-category-service/src/common/helper"
	v "book-category-service/src/common/validator"
	"book-category-service/src/interface/repository"
	"book-category-service/src/interface/service"
	"book-category-service/src/model/dto"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type BookCategoryImpl struct {
	bookCategoryRepo repository.BookCategory
	bookCategoryAggr repository.BookCategoryAggregation
	grpcClient       *client.Grpc
}

func NewBookCategory(br repository.BookCategory, ba repository.BookCategoryAggregation, gc *client.Grpc) service.BookCategory {
	return &BookCategoryImpl{
		bookCategoryRepo: br,
		bookCategoryAggr: ba,
		grpcClient:       gc,
	}
}

func (s *BookCategoryImpl) Create(ctx context.Context, data *dto.CreateBookCategoryReq) error {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	err := s.bookCategoryRepo.Create(ctx, data)
	return err
}

func (s *BookCategoryImpl) FindManyByCategory(ctx context.Context, data *dto.FindManyByCategoryReq) (*dto.DataWithPaging[[]*pb.Book], error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	limit, offset := helper.CreateLimitAndOffset(data.Page)

	totalBooks, err := s.bookCategoryAggr.CountByCategory(ctx, data.Category)
	if err != nil {
		return nil, err
	}

	if totalBooks == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "no books match this category"}
	}

	bookIds, err := s.bookCategoryRepo.FindManyByCategory(ctx, data.Category, limit, offset)
	if err != nil {
		return nil, err
	}

	books, err := s.grpcClient.Book.FindManyByIds(ctx, bookIds)
	if err != nil {
		return nil, err
	}

	return helper.FormatPagedData(books, totalBooks, data.Page, limit), nil
}

func (s *BookCategoryImpl) Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	err := s.bookCategoryRepo.Delete(ctx, data)
	return err
}
