package test

import (
	"book-service/src/interface/service"
	"book-service/src/mock/repository"
	"book-service/src/model/entity"
	serviceimpl "book-service/src/service"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_FindManyPopularBook$ -v ./src/service/test/ -count=1

type FindManyPopularBookTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *FindManyPopularBookTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *FindManyPopularBookTestSuite) Test_Succcess() {
	books := t.createProtobufBooks()
	t.MockPopularTitleKeyRepo_FindManyPopularBook(books, nil)

	res, err := t.bookService.FindManyPopularBook(context.Background())
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}


func (t *FindManyPopularBookTestSuite) MockPopularTitleKeyRepo_FindManyPopularBook(returnArg1 []*entity.Book, returnArg2 error) {

	t.popularTitleKeyRepo.Mock.On("FindManyPopularBook", mock.Anything).Return(returnArg1, returnArg2)
}

func (t *FindManyPopularBookTestSuite) createProtobufBooks() []*entity.Book {
	return []*entity.Book{
		{
			BookId:        1,
			Title:         "The Adventures of Sherlock Holmes",
			Author:        "Arthur Conan Doyle",
			ISBN:          "978-0451524935",
			PublishedYear: 1892,
			Stock:         10,
			Location:      "Aisle 3, Shelf 5",
			CreatedAt:     time.Now().AddDate(-2, 0, 0),
			UpdatedAt:     time.Now(),
		},
		{
			BookId:        2,
			Title:         "1984",
			Author:        "George Orwell",
			ISBN:          "978-0451524936",
			PublishedYear: 1949,
			Stock:         5,
			Location:      "Aisle 1, Shelf 2",
			CreatedAt:     time.Now().AddDate(-1, 0, 0),
			UpdatedAt:     time.Now(),
		},
		{
			BookId:        3,
			Title:         "To Kill a Mockingbird",
			Author:        "Harper Lee",
			ISBN:          "978-0451524937",
			PublishedYear: 1960,
			Stock:         8,
			Location:      "Aisle 2, Shelf 3",
			CreatedAt:     time.Now().AddDate(-3, 0, 0),
			UpdatedAt:     time.Now(),
		},
		{
			BookId:        4,
			Title:         "The Great Gatsby",
			Author:        "F. Scott Fitzgerald",
			ISBN:          "978-0451524938",
			PublishedYear: 1925,
			Stock:         15,
			Location:      "Aisle 4, Shelf 1",
			CreatedAt:     time.Now().AddDate(-5, 0, 0),
			UpdatedAt:     time.Now(),
		},
		{
			BookId:        5,
			Title:         "Moby Dick",
			Author:        "Herman Melville",
			ISBN:          "978-0451524939",
			PublishedYear: 1851,
			Stock:         7,
			Location:      "Aisle 5, Shelf 4",
			CreatedAt:     time.Now().AddDate(-10, 0, 0),
			UpdatedAt:     time.Now(),
		},
	}
}

func TestService_FindManyPopularBook(t *testing.T) {
	suite.Run(t, new(FindManyPopularBookTestSuite))
}