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

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_CreateBookCategory$ -v ./src/service/test/ -count=1

type CreateBookCategoryTestSuite struct {
	suite.Suite
	bookCategoryService      service.BookCategory
	bookCategoryGrpcDelivery *delivery.BookGrpcMock
	bookCategoryRepo         *repository.BookCategoryMock
	bookCategoryAggr         *repository.BookCategoryAggregationMock
}

func (t *CreateBookCategoryTestSuite) SetupSuite() {
	t.bookCategoryRepo = repository.NewBookCategoryMock()
	t.bookCategoryAggr = repository.NewBookCategoryAggregationMock()

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryAggregationMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := grpclient.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)

	t.bookCategoryService = serviceimpl.NewBookCategory(t.bookCategoryRepo, t.bookCategoryAggr, grpcClient)
}

func (t *CreateBookCategoryTestSuite) Test_Succcess() {
	req := t.createBookCategoryReq()

	t.bookCategoryRepo.Mock.On("Create", mock.Anything, req).Return(nil)

	err := t.bookCategoryService.Create(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *CreateBookCategoryTestSuite) Test_WithoutCategory() {
	req := t.createBookCategoryReq()
	req.Category = ""

	err := t.bookCategoryService.Create(context.Background(), req)
	assert.Error(t.T(), err)

	errVldtn, ok := err.(validator.ValidationErrors)
	assert.True(t.T(), ok)
	assert.Equal(t.T(), "Category",errVldtn[0].Field())
}

func (t *CreateBookCategoryTestSuite) createBookCategoryReq() *dto.CreateBookCategoryReq {
	return &dto.CreateBookCategoryReq{
		Category: entity.ADVENTURE,
		BookIds:  []int{1, 2, 3, 4, 5},
	}
}

func TestService_CreateBookCategory(t *testing.T) {
	suite.Run(t, new(CreateBookCategoryTestSuite))
}
