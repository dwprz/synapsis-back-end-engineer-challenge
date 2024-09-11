package repository

import (
	"book-service/src/model/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

type PopularTitleKey interface {
	Upsert(ctx context.Context, title_key string)
	FindManyPopularBook(ctx context.Context) ([]*entity.Book, error)
}

type PopularTitleKeyMock struct {
	mock.Mock
}

func NewPopularTitleKeyMock() *PopularTitleKeyMock {
	return &PopularTitleKeyMock{
		Mock: mock.Mock{},
	}
}

func (d *PopularTitleKeyMock) Upsert(ctx context.Context, title_key string) {}

func (d *PopularTitleKeyMock) FindManyPopularBook(ctx context.Context) ([]*entity.Book, error) {
	arguments := d.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*entity.Book), arguments.Error(1)
}
