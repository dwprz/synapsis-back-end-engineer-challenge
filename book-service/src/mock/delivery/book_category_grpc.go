package delivery

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type BookCategoryGrpcMock struct {
	mock.Mock
}

func NewBookCategoryGrpcMock() *BookCategoryGrpcMock {
	return &BookCategoryGrpcMock{
		Mock: mock.Mock{},
	}
}

func (d *BookCategoryGrpcMock) DeleteBookFromCategoryReq(ctx context.Context, bookId int) error {
	argument := d.Mock.Called(ctx, bookId)

	return argument.Error(0)
}
