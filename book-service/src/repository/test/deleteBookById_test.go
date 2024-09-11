package test

import (
	"book-service/src/api/grpc/client"
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	"book-service/src/mock/delivery"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_DeleteBookById$ -v ./src/repository/test/ -count=1

type DeleteBookByIdTestSuite struct {
	suite.Suite
	postgresDB               *pgxpool.Pool
	bookCategoryGrpcDelivery *delivery.BookCategoryGrpcMock
	bookRepo                 repository.Book
	bookCategoryTestUtil     *util.BookTest
}

func (t *DeleteBookByIdTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryGrpcMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)
	t.bookRepo = repoimpl.NewBook(t.postgresDB, grpcClient)

	t.bookCategoryTestUtil.CreateMany()
}

func (t *DeleteBookByIdTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *DeleteBookByIdTestSuite) Test_Success() {
	bookId := 1
	t.MockBookCategoryGrpcDelivery_DeleteBookFromCategoryReq(bookId, nil)

	err := t.bookRepo.DeleteById(context.Background(), bookId)
	assert.NoError(t.T(), err)
}

func (t *DeleteBookByIdTestSuite) MockBookCategoryGrpcDelivery_DeleteBookFromCategoryReq(bookId int, returnArg1 error) {

	t.bookCategoryGrpcDelivery.Mock.On("DeleteBookFromCategoryReq", mock.Anything, bookId).Return(returnArg1)
}

func TestRepository_DeleteBookById(t *testing.T) {
	suite.Run(t, new(DeleteBookByIdTestSuite))
}
