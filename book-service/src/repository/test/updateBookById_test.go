package test

import (
	"book-service/src/api/grpc/client"
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	"book-service/src/mock/delivery"
	"book-service/src/model/dto"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_UpdateBookById$ -v ./src/repository/test/ -count=1

type UpdateBookByIdTestSuite struct {
	suite.Suite
	postgresDB               *pgxpool.Pool
	bookCategoryGrpcDelivery *delivery.BookCategoryGrpcMock
	bookRepo                 repository.Book
	bookCategoryTestUtil     *util.BookTest
}

func (t *UpdateBookByIdTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryGrpcMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)
	t.bookRepo = repoimpl.NewBook(t.postgresDB, grpcClient)

	t.bookCategoryTestUtil.CreateMany()
}

func (t *UpdateBookByIdTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *UpdateBookByIdTestSuite) Test_UpdateTitle() {
	newTitle := "New Title 1"
	req := t.createUpdateBookReq()
	req.Title = newTitle

	res, err := t.bookRepo.UpdateById(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), newTitle, res.Title)
}

func (t *UpdateBookByIdTestSuite) Test_UpdateStock() {
	newStock := 200
	req := t.createUpdateBookReq()
	req.Stock = &newStock

	res, err := t.bookRepo.UpdateById(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), newStock, res.Stock)
}

func (t *UpdateBookByIdTestSuite) createUpdateBookReq() *dto.UpdateBookReq {
	return &dto.UpdateBookReq{
		BookId: 1,
	}
}

func TestRepository_UpdateBookById(t *testing.T) {
	suite.Run(t, new(UpdateBookByIdTestSuite))
}
