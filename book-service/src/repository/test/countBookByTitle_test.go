package test

import (
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CountBookByTitle$ -v ./src/repository/test/ -count=1

type CountBookByTitleTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookAggrRepo         repository.BookAggregation
	bookCategoryTestUtil *util.BookTest
}

func (t *CountBookByTitleTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.bookAggrRepo = repoimpl.NewBookAggregation(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *CountBookByTitleTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *CountBookByTitleTestSuite) Test_Success() {
	title := "amazing adventure"

	res, err := t.bookAggrRepo.CountByTitle(context.Background(), title)
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), 1, res)
}

func TestRepository_CountBookByTitle(t *testing.T) {
	suite.Run(t, new(CountBookByTitleTestSuite))
}
