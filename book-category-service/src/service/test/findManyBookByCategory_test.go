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

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_FindManyByCategory$ -v ./src/service/test/ -count=1

type FindManyByCategoryTestSuite struct {
	suite.Suite
	bookCategoryService      service.BookCategory
	bookCategoryGrpcDelivery *delivery.BookGrpcMock
	bookCategoryRepo         *repository.BookCategoryMock
	bookCategoryAggr         *repository.BookCategoryAggregationMock
}

func (t *FindManyByCategoryTestSuite) SetupSuite() {
	t.bookCategoryRepo = repository.NewBookCategoryMock()
	t.bookCategoryAggr = repository.NewBookCategoryAggregationMock()

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryAggregationMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := grpclient.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)

	t.bookCategoryService = serviceimpl.NewBookCategory(t.bookCategoryRepo, t.bookCategoryAggr, grpcClient)
}

func (t *FindManyByCategoryTestSuite) Test_Succcess() {
	req := t.createFindManyByCategoryReq()

	bookIds := []int{1, 2, 3, 4, 5}

	t.MockBookCategoryAggr_CountByCategory(req.Category, 5, nil)
	t.MockBookCategoryRepo_FindManyByCategory(req.Category, bookIds, nil)

	books := t.createProtobufBooks()
	t.MockBookCategoryGrpcDelivery_FindManyByIds(bookIds, books, nil)

	res, err := t.bookCategoryService.FindManyByCategory(context.Background(), req)
	assert.NoError(t.T(), err)
	assert.NotEmpty(t.T(), res)
}

func (t *FindManyByCategoryTestSuite) Test_WithoutCategory() {
	req := t.createFindManyByCategoryReq()
	req.Category = ""

	res, err := t.bookCategoryService.FindManyByCategory(context.Background(), req)
	assert.Error(t.T(), err)

	errVldtn, ok := err.(validator.ValidationErrors)
	assert.True(t.T(), ok)
	assert.Equal(t.T(), "Category",errVldtn[0].Field())

	assert.Nil(t.T(), res)
}

func (t *FindManyByCategoryTestSuite) createFindManyByCategoryReq() *dto.FindManyByCategoryReq {
	return &dto.FindManyByCategoryReq{
		Category: entity.ADVENTURE,
		Page:     1,
	}
}

func (t *FindManyByCategoryTestSuite) createProtobufBooks() []*pb.Book {
	return []*pb.Book{
		{
			BookId:        1,
			Title:         "The Adventure of Sherlock Holmes",
			Author:        "Arthur Conan Doyle",
			Isbn:          "978-0-123456-47-2",
			Synopsis:      "A collection of twelve short stories featuring Sherlock Holmes.",
			PublishedYear: 1892,
			Stock:         10,
			Location:      "Shelf A3",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        2,
			Title:         "1984",
			Author:        "George Orwell",
			Isbn:          "978-0-987654-32-1",
			Synopsis:      "A dystopian novel set in a totalitarian society ruled by Big Brother.",
			PublishedYear: 1949,
			Stock:         8,
			Location:      "Shelf B1",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        3,
			Title:         "To Kill a Mockingbird",
			Author:        "Harper Lee",
			Isbn:          "978-0-112233-44-5",
			Synopsis:      "A novel about the serious issues of rape and racial inequality.",
			PublishedYear: 1960,
			Stock:         15,
			Location:      "Shelf C5",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        4,
			Title:         "Pride and Prejudice",
			Author:        "Jane Austen",
			Isbn:          "978-0-998877-11-6",
			Synopsis:      "A romantic novel of manners.",
			PublishedYear: 1813,
			Stock:         5,
			Location:      "Shelf D2",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        5,
			Title:         "The Catcher in the Rye",
			Author:        "J.D. Salinger",
			Isbn:          "978-0-776655-88-9",
			Synopsis:      "A story about teenage rebellion and angst.",
			PublishedYear: 1951,
			Stock:         7,
			Location:      "Shelf E4",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
	}
}

func (t *FindManyByCategoryTestSuite) MockBookCategoryAggr_CountByCategory(data entity.Category, resturnArg1 int, resturnArg2 error) {

	t.bookCategoryAggr.Mock.On("CountByCategory", mock.Anything, data, mock.Anything, mock.Anything).Return(resturnArg1, resturnArg2)
}

func (t *FindManyByCategoryTestSuite) MockBookCategoryRepo_FindManyByCategory(data entity.Category, resturnArg1 []int, resturnArg2 error) {

	t.bookCategoryRepo.Mock.On("FindManyByCategory", mock.Anything, data, mock.Anything, mock.Anything).Return(resturnArg1, resturnArg2)
}

func (t *FindManyByCategoryTestSuite) MockBookCategoryGrpcDelivery_FindManyByIds(bookIds []int, resturnArg1 []*pb.Book, resturnArg2 error) {

	t.bookCategoryGrpcDelivery.Mock.On("FindManyByIds", mock.Anything, bookIds).Return(resturnArg1, resturnArg2)
}

func TestService_FindManyByCategory(t *testing.T) {
	suite.Run(t, new(FindManyByCategoryTestSuite))
}
