package test

import (
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	"book-service/src/model/dto"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CountBookByFields$ -v ./src/repository/test/ -count=1

type CountBookByFieldsTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookAggrRepo         repository.BookAggregation
	bookCategoryTestUtil *util.BookTest
}

func (t *CountBookByFieldsTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookAggrRepo = repoimpl.NewBookAggregation(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *CountBookByFieldsTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *CountBookByFieldsTestSuite) Test_Success() {
	author := "Arthur Conan Doyle"

	res, err := t.bookAggrRepo.CountByFields(context.Background(), &dto.GetBookReq{Author: author})
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), 1, res)
}

func TestRepository_CountBookByFields(t *testing.T) {
	suite.Run(t, new(CountBookByFieldsTestSuite))
}
