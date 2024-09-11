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
// go test -run ^TestRepository_CreateBookCategory$ -v ./src/repository/test/ -count=1

type CreateBookCategoryTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	bookCategoryRepo     repository.BookCategory
	bookCategoryTestUtil *util.BookCategoryTest
}

func (t *CreateBookCategoryTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookCategoryTest(t.postgresDB)

	t.bookCategoryRepo = repoimpl.NewBookCategory(t.postgresDB)
}

func (t *CreateBookCategoryTestSuite) TearDownTest() {
	t.bookCategoryTestUtil.Delete()
}

func (t *CreateBookCategoryTestSuite) TearDownSuite() {
	t.postgresDB.Close()
}

func (t *CreateBookCategoryTestSuite) Test_Succcess() {
	err := t.bookCategoryRepo.Create(context.Background(), &dto.CreateBookCategoryReq{
		Category: entity.ADVENTURE,
		BookIds:  []int{1, 2, 3, 4, 5},
	})

	assert.NoError(t.T(), err)
}

func (t *CreateBookCategoryTestSuite) Test_InvalidCategory() {
	err := t.bookCategoryRepo.Create(context.Background(), &dto.CreateBookCategoryReq{
		Category: "invalid category",
		BookIds:  []int{1, 2, 3, 4, 5},
	})

	assert.Error(t.T(), err)
}

func TestRepository_CreateBookCategory(t *testing.T) {
	suite.Run(t, new(CreateBookCategoryTestSuite))
}
