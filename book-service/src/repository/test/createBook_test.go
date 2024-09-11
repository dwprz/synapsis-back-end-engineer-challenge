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
// go test -run ^TestRepository_CreateBook$ -v ./src/repository/test/ -count=1

type CreateBookTestSuite struct {
	suite.Suite
	postgresDB               *pgxpool.Pool
	bookCategoryGrpcDelivery *delivery.BookCategoryGrpcMock
	bookRepo                 repository.Book
	bookCategoryTestUtil     *util.BookTest
}

func (t *CreateBookTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryGrpcMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)

	t.bookRepo = repoimpl.NewBook(t.postgresDB, grpcClient)
}

func (t *CreateBookTestSuite) TearDownTest() {
	t.bookCategoryTestUtil.Delete()
}

func (t *CreateBookTestSuite) TearDownSuite() {
	t.postgresDB.Close()
}

func (t *CreateBookTestSuite) Test_Succcess() {
	req := t.CreateAddBookReq()

	err := t.bookRepo.Create(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *CreateBookTestSuite) CreateAddBookReq() *dto.AddBookReq {
	return &dto.AddBookReq{
		Title:         "Amazing Story",
		Author:        "John Doe",
		ISBN:          "978-0-123456-48-9",
		PublishedYear: 1999,
		Stock:         20,
		Location:      "Shelf B",
	}
}

func TestRepository_CreateBook(t *testing.T) {
	suite.Run(t, new(CreateBookTestSuite))
}
