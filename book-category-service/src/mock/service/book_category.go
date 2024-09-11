package service

import (
	"book-category-service/src/model/dto"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/stretchr/testify/mock"
)

type BookCategoryMock struct {
	mock.Mock
}

func NewBookCategoryMock() *BookCategoryMock {
	return &BookCategoryMock{
		Mock: mock.Mock{},
	}
}

func (s *BookCategoryMock) Create(ctx context.Context, data *dto.CreateBookCategoryReq) error {
	argument := s.Mock.Called(ctx, data)

	return argument.Error(0)
}

func (s *BookCategoryMock) FindManyByCategory(ctx context.Context, data *dto.FindManyByCategoryReq) (*dto.DataWithPaging[[]*pb.Book], error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.DataWithPaging[[]*pb.Book]), arguments.Error(1)
}

func (s *BookCategoryMock) Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error {
	argument := s.Mock.Called(ctx, data)

	return argument.Error(0)
}
