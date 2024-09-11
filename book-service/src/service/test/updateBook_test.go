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
// go test -run ^TestService_UpdateBook$ -v ./src/service/test/ -count=1

type UpdateBookTestSuite struct {
	suite.Suite
	bookService         service.Book
	bookRepo            *repository.BookMock
	bookAggrRepo        *repository.BookAggrMock
	popularTitleKeyRepo *repository.PopularTitleKeyMock
}

func (t *UpdateBookTestSuite) SetupSuite() {
	t.bookRepo = repository.NewBookMock()
	t.bookAggrRepo = repository.NewBookAggrMock()
	t.popularTitleKeyRepo = repository.NewPopularTitleKeyMock()

	t.bookService = serviceimpl.NewBook(t.bookRepo, t.bookAggrRepo, t.popularTitleKeyRepo)
}

func (t *UpdateBookTestSuite) Test_Succcess() {
	req := t.createUpdateBookReq()

	books := t.createEntityBook()

	t.MockBookRepo_UpdateById(req, books, nil)

	res, err := t.bookService.Update(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func (t *UpdateBookTestSuite) Test_WithoutBookId() {
	req := t.createUpdateBookReq()
	req.BookId = 0

	res, err := t.bookService.Update(context.Background(), req)
	assert.Error(t.T(), err)

	assert.Nil(t.T(), res)
}

func (t *UpdateBookTestSuite) MockBookRepo_UpdateById(data *dto.UpdateBookReq, returnArg1 *entity.Book, returnArg2 error) {

	t.bookRepo.Mock.On("UpdateById", mock.Anything, data).Return(returnArg1, returnArg2)
}

func (t *UpdateBookTestSuite) createUpdateBookReq() *dto.UpdateBookReq {
	return &dto.UpdateBookReq{
		BookId: 1,
		Title:  "The Adventures of Sherlock Holmes",
	}
}

func (t *UpdateBookTestSuite) createEntityBook() *entity.Book {
	return &entity.Book{
		BookId:        1,
		Title:         "The Adventures of Sherlock Holmes",
		Author:        "Arthur Conan Doyle",
		ISBN:          "978-0451524935",
		PublishedYear: 1892,
		Stock:         10,
		Location:      "Aisle 3, Shelf 5",
		CreatedAt:     time.Now().AddDate(-2, 0, 0),
		UpdatedAt:     time.Now(),
	}
}

func TestService_UpdateBook(t *testing.T) {
	suite.Run(t, new(UpdateBookTestSuite))
}
