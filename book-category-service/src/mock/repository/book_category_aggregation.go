package repository

import (
	"book-category-service/src/model/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

type BookCategoryAggregationMock struct {
	mock.Mock
}

func NewBookCategoryAggregationMock() *BookCategoryAggregationMock {
	return &BookCategoryAggregationMock{
		Mock: mock.Mock{},
	}
}

func (r *BookCategoryAggregationMock) CountByCategory(ctx context.Context, category entity.Category) (totalBooks int, err error) {
	argument := r.Mock.Called(ctx, category)

	return argument.Int(0), argument.Error(1)
}
