package test

import (
	"context"
	"testing"
	"user-service/src/infrastructure/database"
	"user-service/src/interface/repository"
	"user-service/src/mock/cache"
	repoimpl "user-service/src/repository"
	"user-service/src/test/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_SetNullRefreshToken$ -v ./src/repository/test/ -count=1

type SetNullRefreshTokenTestSuite struct {
	suite.Suite
	userRepo     repository.User
	userCache    *cache.UserMock
	postgresDB   *pgxpool.Pool
	userTestUtil *util.UserTest
}

func (t *SetNullRefreshTokenTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.userCache = cache.NewUserMock()

	t.userTestUtil = util.NewUserTest(t.postgresDB)
	t.userRepo = repoimpl.NewUser(t.postgresDB, t.userCache)

	t.userTestUtil.Create()
}

func (t *SetNullRefreshTokenTestSuite) TearDownSuite() {
	t.userTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *SetNullRefreshTokenTestSuite) Test_Success() {
	refreshToken := "example-refresh-token"
	err := t.userRepo.SetNullRefreshToken(context.Background(), refreshToken)

	assert.NoError(t.T(), err)
}

func TestRepository_SetNullRefreshToken(t *testing.T) {
	suite.Run(t, new(SetNullRefreshTokenTestSuite))
}
