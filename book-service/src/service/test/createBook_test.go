package test

import (
	"book-service/src/interface/service"
	"book-service/src/mock/repository"
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	serviceimpl "book-service/src/service"
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_AddBook$ -v ./src/service/test/ -count=1

type AddBookTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *AddBookTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *AddBookTestSuite) Test_Succcess() {
	req := t.createAddBookReq()

	t.MockBookRepo_FindByTitle(req.Title, nil, nil)
	t.MockBookRepo_Create(req, nil)

	err := t.bookService.Add(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *AddBookTestSuite) Test_WithoutTitle() {
	req := t.createAddBookReq()
	req.Title = ""

	err := t.bookService.Add(context.Background(), req)
	assert.Error(t.T(), err)

	errVldtn, ok := err.(validator.ValidationErrors)
	assert.True(t.T(), ok)

	assert.Equal(t.T(), "Title", errVldtn[0].Field())
}

func (t *AddBookTestSuite) MockBookRepo_FindByTitle(title string, returnArg1 *entity.Book, returnArg2 error) {

	t.bookRepo.Mock.On("FindByTitle", mock.Anything, title).Return(returnArg1, returnArg2)
}

func (t *AddBookTestSuite) MockBookRepo_Create(data *dto.AddBookReq, returnArg1 error) {

	t.bookRepo.Mock.On("Create", mock.Anything, data).Return(returnArg1)
}

func (t *AddBookTestSuite) createAddBookReq() *dto.AddBookReq {
	return &dto.AddBookReq{
		Title:         "The Adventure of Sherlock Holmes",
		Author:        "Arthur Conan Doyle",
		ISBN:          "978-0-123456-47-2",
		PublishedYear: 1892,
		Stock:         10,
		Location:      "Shelf A3",
	}
}

func TestService_AddBook(t *testing.T) {
	suite.Run(t, new(AddBookTestSuite))
}
