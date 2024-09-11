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
// go test -run ^TestRepository_DeleteUser$ -v ./src/repository/test/ -count=1

type deleteUserTestSuite struct {
	suite.Suite
	userRepo     repository.User
	userCache    *cache.UserMock
	postgresDB   *pgxpool.Pool
	userTestUtil *util.UserTest
}

func (t *deleteUserTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.userCache = cache.NewUserMock()

	t.userTestUtil = util.NewUserTest(t.postgresDB)
	t.userRepo = repoimpl.NewUser(t.postgresDB, t.userCache)

	t.userTestUtil.Create()
}

func (t *deleteUserTestSuite) SetupTest() {
	t.userTestUtil.Create()
}

func (t *deleteUserTestSuite) TearDownSuite() {
	t.postgresDB.Close()
}

func (t *deleteUserTestSuite) TearDownTest() {
	t.userTestUtil.Delete()
}

func (t *deleteUserTestSuite) Test_Success() {
	userId := "ynA1nZIULkXLrfy0fvz5t"
	err := t.userRepo.Delete(context.Background(), userId)

	assert.NoError(t.T(), err)
}

func TestRepository_DeleteUser(t *testing.T) {
	suite.Run(t, new(deleteUserTestSuite))
}
