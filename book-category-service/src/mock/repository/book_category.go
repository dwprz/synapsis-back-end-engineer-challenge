package repository

import (
	"book-category-service/src/model/dto"
	"book-category-service/src/model/entity"
	"context"

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

func (r *BookCategoryMock) Create(ctx context.Context, data *dto.CreateBookCategoryReq) error {
	argument := r.Mock.Called(ctx, data)

	return argument.Error(0)
}

func (r *BookCategoryMock) FindManyByCategory(ctx context.Context, category entity.Category, limit, offset int) (bookIds []int, err error) {
	arguments := r.Mock.Called(ctx, category, limit, offset)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]int), arguments.Error(1)
}

func (r *BookCategoryMock) Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error {
	argument := r.Mock.Called(ctx, data)

	return argument.Error(0)
}
