package test

import (
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_FindManyPopularBook$ -v ./src/repository/test/ -count=1

type FindManyPopularBookTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	popularTitleKeyRepo  repository.PopularTitleKey
	bookCategoryTestUtil *util.BookTest
}

func (t *FindManyPopularBookTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.popularTitleKeyRepo = repoimpl.NewPopularTitleKey(t.postgresDB)
	
	t.bookCategoryTestUtil.AddPopularTitleKey()
	t.bookCategoryTestUtil.CreateMany()
}

func (t *FindManyPopularBookTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *FindManyPopularBookTestSuite) Test_Success() {
	res, err := t.popularTitleKeyRepo.FindManyPopularBook(context.Background())
	assert.NoError(t.T(), err)

	fmt.Println(res)
}

func TestRepository_FindManyPopularBook(t *testing.T) {
	suite.Run(t, new(FindManyPopularBookTestSuite))
}
