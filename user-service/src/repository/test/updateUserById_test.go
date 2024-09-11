package test

import (
	"context"
	"testing"
	"user-service/src/infrastructure/database"
	"user-service/src/interface/repository"
	"user-service/src/mock/cache"
	"user-service/src/model/dto"
	repoimpl "user-service/src/repository"
	"user-service/src/test/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_UpdateUserById$ -v ./src/repository/test/ -count=1

type UpdateUserByIdTestSuite struct {
	suite.Suite
	userRepo     repository.User
	userCache    *cache.UserMock
	postgresDB   *pgxpool.Pool
	userTestUtil *util.UserTest
}

func (t *UpdateUserByIdTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.userCache = cache.NewUserMock()

	t.userTestUtil = util.NewUserTest(t.postgresDB)
	t.userRepo = repoimpl.NewUser(t.postgresDB, t.userCache)

	t.userTestUtil.Create()
}

func (t *UpdateUserByIdTestSuite) TearDownSuite() {
	t.userTestUtil.Delete()
	t.postgresDB.Close()
}

func (t *UpdateUserByIdTestSuite) Test_Success() {

	res, err := t.userRepo.UpdateByUserId(context.Background(), &dto.UpdateUserReq{
		UserId:   "ynA1nZIULkXLrfy0fvz5t",
		FullName: "New John Doe",
	})

	assert.NoError(t.T(), err)
	assert.NotEmpty(t.T(), res)
}

func TestRepository_UpdateUserById(t *testing.T) {
	suite.Run(t, new(UpdateUserByIdTestSuite))
}
