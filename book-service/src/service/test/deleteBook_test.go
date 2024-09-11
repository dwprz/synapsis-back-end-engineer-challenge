package test

import (
	"book-service/src/interface/service"
	"book-service/src/mock/repository"
	serviceimpl "book-service/src/service"
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_DeleteBook$ -v ./src/service/test/ -count=1

type DeleteBookTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *DeleteBookTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *DeleteBookTestSuite) Test_Succcess() {
	bookId := 1
	t.MockBookRepo_DeleteById(bookId, nil)

	err := t.bookService.Delete(context.Background(), bookId)
	assert.NoError(t.T(), err)
}

func (t *DeleteBookTestSuite) Test_WithoutBookId() {

	err := t.bookService.Delete(context.Background(), 0)
	assert.Error(t.T(), err)

	errVldtn, ok := err.(validator.ValidationErrors)
	assert.True(t.T(), ok)

	assert.Equal(t.T(), "Id", errVldtn[0].Field())
}

func (t *DeleteBookTestSuite) MockBookRepo_DeleteById(bookId int, returnArg1 error) {

	t.bookRepo.Mock.On("DeleteById", mock.Anything, bookId).Return(returnArg1)
}

func TestService_DeleteBook(t *testing.T) {
	suite.Run(t, new(DeleteBookTestSuite))
}
