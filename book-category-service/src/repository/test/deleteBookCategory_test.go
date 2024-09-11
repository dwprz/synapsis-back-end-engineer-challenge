package test

import (
	"book-category-service/src/infrastructure/database"
	"book-category-service/src/interface/repository"
	"book-category-service/src/model/dto"
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
// go test -run ^TestRepository_DeleteBookCategory$ -v ./src/repository/test/ -count=1

type DeleteBookCategoryTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookCategoryRepo     repository.BookCategory
	bookCategoryTestUtil *util.BookCategoryTest
}

func (t *DeleteBookCategoryTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookCategoryTest(t.postgresDB)

	t.bookCategoryRepo = repoimpl.NewBookCategory(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *DeleteBookCategoryTestSuite) SetupTest() {
	t.bookCategoryTestUtil.CreateMany()
}

func (t *DeleteBookCategoryTestSuite) TearDownTest() {
	t.bookCategoryTestUtil.Delete()
}

func (t *DeleteBookCategoryTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *DeleteBookCategoryTestSuite) Test_Succcess() {
	category := entity.ADVENTURE

	err := t.bookCategoryRepo.Delete(context.Background(), &dto.DeleteBookCategoryReq{
		Category: &category,
	})

	assert.NoError(t.T(), err)
}

func (t *DeleteBookCategoryTestSuite) Test_NilArguments() {
	err := t.bookCategoryRepo.Delete(context.Background(), &dto.DeleteBookCategoryReq{})
	assert.Error(t.T(), err)
}

func TestRepository_DeleteBookCategory(t *testing.T) {
	suite.Run(t, new(DeleteBookCategoryTestSuite))
}
