package test

import (
	grpclient "book-category-service/src/api/grpc/client"
	"book-category-service/src/interface/service"
	"book-category-service/src/mock/delivery"
	"book-category-service/src/mock/repository"
	"book-category-service/src/model/dto"
	"book-category-service/src/model/entity"
	serviceimpl "book-category-service/src/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_DeleteCategory$ -v ./src/service/test/ -count=1

type DeleteBookCategoryTestSuite struct {
	suite.Suite
	bookCategoryService      service.BookCategory
	bookCategoryGrpcDelivery *delivery.BookGrpcMock
	bookCategoryRepo         *repository.BookCategoryMock
	bookCategoryAggr         *repository.BookCategoryAggregationMock
}

func (t *DeleteBookCategoryTestSuite) SetupSuite() {
	t.bookCategoryRepo = repository.NewBookCategoryMock()
	t.bookCategoryAggr = repository.NewBookCategoryAggregationMock()

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryAggregationMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := grpclient.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)

	t.bookCategoryService = serviceimpl.NewBookCategory(t.bookCategoryRepo, t.bookCategoryAggr, grpcClient)
}

func (t *DeleteBookCategoryTestSuite) Test_Succcess() {
	req := t.createDeleteBookCategoryReq()

	t.MockBookCategoryRepo_DeleteBookCategory(req, nil)

	err := t.bookCategoryService.Delete(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *DeleteBookCategoryTestSuite) Test_WithoutCategory() {
	req := t.createDeleteBookCategoryReq()
	req.Category = nil

	t.MockBookCategoryRepo_DeleteBookCategory(req, nil)

	err := t.bookCategoryService.Delete(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *DeleteBookCategoryTestSuite) createDeleteBookCategoryReq() *dto.DeleteBookCategoryReq {
	category := entity.ADVENTURE
	bookId := 1

	return &dto.DeleteBookCategoryReq{
		Category: &category,
		BookId:   &bookId,
	}
}

func (t *DeleteBookCategoryTestSuite) MockBookCategoryRepo_DeleteBookCategory(data *dto.DeleteBookCategoryReq, resturnArg1 error) {

	t.bookCategoryRepo.Mock.On("Delete", mock.Anything, data).Return(resturnArg1)
}

func TestService_DeleteCategory(t *testing.T) {
	suite.Run(t, new(DeleteBookCategoryTestSuite))
}
