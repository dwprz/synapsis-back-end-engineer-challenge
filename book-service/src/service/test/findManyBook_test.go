package test

import (
	"book-service/src/interface/service"
	"book-service/src/mock/repository"
	"book-service/src/model/dto"
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
// go test -run ^TestService_FindManyBook$ -v ./src/service/test/ -count=1

type FindManyBookTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *FindManyBookTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *FindManyBookTestSuite) Test_Succcess() {
	req := t.createGetBookReq()
	req.Title = "Example Title"

	books := t.createEntityBook()

	t.MockBookAggrRepo_CountByTitle(req.Title, 20, nil)
	t.MockBookRepo_FindManyByTitle(req.Title, books, nil)

	res, err := t.bookService.FindMany(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func (t *FindManyBookTestSuite) Test_WithoutArgs() {
	err := t.bookService.Add(context.Background(), nil)

	assert.Error(t.T(), err)
}

func (t *FindManyBookTestSuite) MockBookAggrRepo_CountByTitle(title string, returnArg1 int, returnArg2 error) {

	t.bookAggrRepo.Mock.On("CountByTitle", mock.Anything, title).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookAggrRepo_CountByTiltle(title string, returnArg1 int, returnArg2 error) {

	t.bookAggrRepo.Mock.On("CountByTitle", mock.Anything, title).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookRepo_FindManyByTitle(title string, returnArg1 []*entity.Book, returnArg2 error) {

	t.bookRepo.Mock.On("FindManyByTitle", mock.Anything, title, mock.Anything, mock.Anything).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookRepo_FindManyByFields(data *dto.GetBookReq, returnArg1 []*entity.Book, returnArg2 error) {

	t.bookRepo.Mock.On("FindManyByFields", mock.Anything, data, mock.Anything, mock.Anything).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookAggrRepo_CountByFields(data *dto.GetBookReq, returnArg1 int, returnArg2 error) {

	t.bookAggrRepo.Mock.On("CountByFields", mock.Anything, data).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookAggrRepo_CountAll(returnArg1 int, returnArg2 error) {

	t.bookAggrRepo.Mock.On("CountAll", mock.Anything).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) MockBookRepo_FindManyRandom(returnArg1 []*entity.Book, returnArg2 error) {

	t.bookRepo.Mock.On("CountByFields", mock.Anything, mock.Anything, mock.Anything).Return(returnArg1, returnArg2)
}

func (t *FindManyBookTestSuite) createGetBookReq() *dto.GetBookReq {
	return &dto.GetBookReq{
		Page:   1,
		Author: "Arthur Conan Doyle",
	}
}

func (t *FindManyBookTestSuite) createEntityBook() []*entity.Book {
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

func TestService_FindManyBook(t *testing.T) {
	suite.Run(t, new(FindManyBookTestSuite))
}
