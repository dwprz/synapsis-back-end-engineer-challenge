package service

import (
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/stretchr/testify/mock"
)

type BookMock struct {
	mock.Mock
}

func NewBookMock() *BookMock {
	return &BookMock{
		Mock: mock.Mock{},
	}
}

func (s *BookMock) Add(ctx context.Context, data *dto.AddBookReq) error {
	argument := s.Mock.Called(ctx, data)

	return argument.Error(0)
}

func (s *BookMock) FindMany(ctx context.Context, data *dto.GetBookReq) (*dto.DataWithPaging[[]*entity.Book], error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.DataWithPaging[[]*entity.Book]), arguments.Error(1)
}

func (s *BookMock) FindManyByIds(ctx context.Context, data *dto.FindManyByIdsReq) ([]*pb.Book, error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*pb.Book), arguments.Error(1)
}

func (s *BookMock) FindManyPopularBook(ctx context.Context) ([]*entity.Book, error) {
	arguments := s.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*entity.Book), arguments.Error(1)
}

func (s *BookMock) Update(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.Book), arguments.Error(1)
}

func (s *BookMock) Delete(ctx context.Context, bookId int) error {
	argument := s.Mock.Called(ctx, bookId)

	return argument.Error(0)
}
