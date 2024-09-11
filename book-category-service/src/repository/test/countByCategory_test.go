package test

import (
	"book-category-service/src/infrastructure/database"
	"book-category-service/src/interface/repository"
	"book-category-service/src/model/entity"
	repoimpl "book-category-service/src/repository"
	"book-category-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CountByCategory$ -v ./src/repository/test/ -count=1

type CountByCategoryTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookCategoryAggr     repository.BookCategoryAggregation
	bookCategoryTestUtil *util.BookCategoryTest
}

func (t *CountByCategoryTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookCategoryTest(t.postgresDB)

	t.bookCategoryAggr = repoimpl.NewBookCategoryAggregation(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *CountByCategoryTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *CountByCategoryTestSuite) Test_Succcess() {
	res, err := t.bookCategoryAggr.CountByCategory(context.Background(), entity.ADVENTURE)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func (t *CountByCategoryTestSuite) Test_NotFound() {
	res, err := t.bookCategoryAggr.CountByCategory(context.Background(), "not found category")
	assert.Error(t.T(), err)

	assert.Equal(t.T(), 0, res)
}

func TestRepository_CountByCategory(t *testing.T) {
	suite.Run(t, new(CountByCategoryTestSuite))
}
