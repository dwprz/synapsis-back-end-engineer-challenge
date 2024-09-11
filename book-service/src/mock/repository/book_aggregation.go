package repository

import (
	"book-service/src/model/dto"
	"context"

	"github.com/stretchr/testify/mock"
)

type BookAggrMock struct {
	mock.Mock
}

func NewBookAggrMock() *BookAggrMock {
	return &BookAggrMock{
		Mock: mock.Mock{},
	}
}

func (d *BookAggrMock) CountByTitle(ctx context.Context, title string) (totalBooks int, err error) {
	argument := d.Mock.Called(ctx, title)

	return argument.Int(0), argument.Error(1)
}

func (d *BookAggrMock) CountByFields(ctx context.Context, data *dto.GetBookReq) (totalBooks int, err error) {
	argument := d.Mock.Called(ctx, data)

	return argument.Int(0), argument.Error(1)
}

func (d *BookAggrMock) CountAll(ctx context.Context) (totalBooks int, err error) {
	argument := d.Mock.Called(ctx)

	return argument.Int(0), argument.Error(1)
}