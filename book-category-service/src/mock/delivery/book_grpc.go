package delivery

import (
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/stretchr/testify/mock"
)

type BookGrpcMock struct {
	mock.Mock
}

func NewBookCategoryAggregationMock() *BookGrpcMock {
	return &BookGrpcMock{
		Mock: mock.Mock{},
	}
}
func (d *BookGrpcMock) FindManyByIds(ctx context.Context, bookIds []int) ([]*pb.Book, error) {
	arguments := d.Mock.Called(ctx, bookIds)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).([]*pb.Book), arguments.Error(1)
}