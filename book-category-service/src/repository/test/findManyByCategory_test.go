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
// go test -run ^TestRepository_FindManyByCategory$ -v ./src/repository/test/ -count=1

type FindManyByCategoryTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookCategoryRepo     repository.BookCategory
	bookCategoryTestUtil *util.BookCategoryTest
}

func (t *FindManyByCategoryTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookCategoryTest(t.postgresDB)

	t.bookCategoryRepo = repoimpl.NewBookCategory(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *FindManyByCategoryTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *FindManyByCategoryTestSuite) Test_Succcess() {
	res, err := t.bookCategoryRepo.FindManyByCategory(context.Background(), entity.ADVENTURE, 20, 0)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res)
}

func (t *FindManyByCategoryTestSuite) Test_InvalidCategory() {
	res, err := t.bookCategoryRepo.FindManyByCategory(context.Background(), "invalid category", 20, 0)
	assert.Error(t.T(), err)

	assert.Nil(t.T(), res)
}

func TestRepository_FindManyByCategory(t *testing.T) {
	suite.Run(t, new(FindManyByCategoryTestSuite))
}
