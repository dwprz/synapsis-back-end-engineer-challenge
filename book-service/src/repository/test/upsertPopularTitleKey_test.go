package test

import (
	"book-service/src/infrastructure/database"
	"book-service/src/interface/repository"
	repoimpl "book-service/src/repository"
	"book-service/src/test/util"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_UpsertPopularTitleKey$ -v ./src/repository/test/ -count=1

type UpsertPopularTitleKeyTestSuite struct {
	suite.Suite
	postgresDB           *pgxpool.Pool
	popularTitleKeyRepo  repository.PopularTitleKey
	bookCategoryTestUtil *util.BookTest
}

func (t *UpsertPopularTitleKeyTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.bookCategoryTestUtil = util.NewBookTest(t.postgresDB)

	t.popularTitleKeyRepo = repoimpl.NewPopularTitleKey(t.postgresDB)
	t.bookCategoryTestUtil.CreateMany()
}

func (t *UpsertPopularTitleKeyTestSuite) TearDownSuite() {
	t.bookCategoryTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *UpsertPopularTitleKeyTestSuite) Test_Success() {
	t.popularTitleKeyRepo.Upsert(context.Background(), "Amazing Adventure")
}

func TestRepository_UpsertPopularTitleKey(t *testing.T) {
	suite.Run(t, new(UpsertPopularTitleKeyTestSuite))
}
