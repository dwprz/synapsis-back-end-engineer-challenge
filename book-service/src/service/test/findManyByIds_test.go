package test

import (
	"book-service/src/interface/service"
	"book-service/src/mock/repository"
	"book-service/src/model/dto"
	serviceimpl "book-service/src/service"
	"context"
	"testing"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_FindManyByIds$ -v ./src/service/test/ -count=1

type FindManyByIdsTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *FindManyByIdsTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *FindManyByIdsTestSuite) Test_Succcess() {
	req := t.createFindManyByIdsReq()

	books := t.createProtobufBooks()
	t.MockBookRepo_FindManyByIds(req.BookIds, books, nil)

	res, err := t.bookService.FindManyByIds(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func (t *FindManyByIdsTestSuite) Test_NilArguments() {
	res, err := t.bookService.FindManyByIds(context.Background(), nil)
	assert.Error(t.T(), err)

	assert.Nil(t.T(), res)
}

func (t *FindManyByIdsTestSuite) MockBookRepo_FindManyByIds(bookIds []uint32, returnArg1 []*pb.Book, returnArg2 error) {

	t.bookRepo.Mock.On("FindManyByIds", mock.Anything, bookIds).Return(returnArg1, returnArg2)
}

func (t *FindManyByIdsTestSuite) createFindManyByIdsReq() *dto.FindManyByIdsReq {
	return &dto.FindManyByIdsReq{
		BookIds: []uint32{1, 2, 3, 4, 5},
	}
}

func (t *FindManyByIdsTestSuite) createProtobufBooks() []*pb.Book {
	return []*pb.Book{
		{
			BookId:        1,
			Title:         "The Adventures of Sherlock Holmes",
			Author:        "Arthur Conan Doyle",
			Isbn:          "978-0451524935",
			PublishedYear: 1892,
			Stock:         10,
			Location:      "Aisle 3, Shelf 5",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        2,
			Title:         "1984",
			Author:        "George Orwell",
			Isbn:          "978-0451524936",
			PublishedYear: 1949,
			Stock:         5,
			Location:      "Aisle 1, Shelf 2",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        3,
			Title:         "To Kill a Mockingbird",
			Author:        "Harper Lee",
			Isbn:          "978-0451524937",
			PublishedYear: 1960,
			Stock:         8,
			Location:      "Aisle 2, Shelf 3",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        4,
			Title:         "The Great Gatsby",
			Author:        "F. Scott Fitzgerald",
			Isbn:          "978-0451524938",
			PublishedYear: 1925,
			Stock:         15,
			Location:      "Aisle 4, Shelf 1",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
		{
			BookId:        5,
			Title:         "Moby Dick",
			Author:        "Herman Melville",
			Isbn:          "978-0451524939",
			PublishedYear: 1851,
			Stock:         7,
			Location:      "Aisle 5, Shelf 4",
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		},
	}
}

func TestService_FindManyByIds(t *testing.T) {
	suite.Run(t, new(FindManyByIdsTestSuite))
}
