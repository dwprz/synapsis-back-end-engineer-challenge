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
// go test -run ^TestRepository_FindManyBookByFields$ -v ./src/repository/test/ -count=1

type FindManyBookByFieldsTestSuite struct {
	suite.Suite
	postgresDB               *pgxpool.Pool
	bookCategoryGrpcDelivery *delivery.BookCategoryGrpcMock
	bookRepo                 repository.Book
	bookCategoryTestUtil     *util.BookTest
}

func (t *FindManyBookByFieldsTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookCategoryGrpcDelivery = delivery.NewBookCategoryGrpcMock()
	bookCategoryGrpcConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(t.bookCategoryGrpcDelivery, bookCategoryGrpcConn)
	t.bookRepo = repoimpl.NewBook(t.postgresDB, grpcClient)

	t.bookCategoryTestUtil.CreateMany()
}

func (t *FindManyBookByFieldsTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *FindManyBookByFieldsTestSuite) Test_Succcess() {
	author := "Arthur Conan Doyle"

	res, err := t.bookRepo.FindManyByFields(context.Background(), &dto.GetBookReq{Author: author}, 20, 0)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func TestRepository_FindManyBookByFields(t *testing.T) {
	suite.Run(t, new(FindManyBookByFieldsTestSuite))
}
