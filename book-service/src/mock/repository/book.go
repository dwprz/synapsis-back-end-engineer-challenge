package repository

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

func (d *BookMock) Create(ctx context.Context, data *dto.AddBookReq) error {
	argument := d.Mock.Called(ctx, data)

	return argument.Error(0)
}

func (d *BookMock) FindByTitle(ctx context.Context, title string) (*entity.Book, error) {
	arguments := d.Mock.Called(ctx, title)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.Book), arguments.Error(1)
}

func (d *BookMock) FindManyByTitle(ctx context.Context, title string, limit, offset int) ([]*entity.Book, error) {
	arguments := d.Mock.Called(ctx, title)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*entity.Book), arguments.Error(1)
}

func (d *BookMock) FindManyByFields(ctx context.Context, data *dto.GetBookReq, limit, offset int) ([]*entity.Book, error) {
	arguments := d.Mock.Called(ctx, data, limit, offset)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*entity.Book), arguments.Error(1)
}

func (d *BookMock) FindManyByRandom(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	arguments := d.Mock.Called(ctx, limit, offset)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*entity.Book), arguments.Error(1)
}

func (d *BookMock) FindManyByIds(ctx context.Context, ids []uint32) ([]*pb.Book, error) {
	arguments := d.Mock.Called(ctx, ids)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*pb.Book), arguments.Error(1)
}

func (d *BookMock) UpdateById(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error) {
	arguments := d.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.Book), arguments.Error(1)
}

func (d *BookMock) DeleteById(ctx context.Context, bookId int) error {
	arguments := d.Mock.Called(ctx, bookId)

	if arguments.Get(0) == nil {
		return arguments.Error(0)
	}

	return arguments.Error(0)
}
