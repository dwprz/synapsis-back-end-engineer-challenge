package test

import (
	"context"
	"testing"
	"user-service/src/common/log"
	"user-service/src/infrastructure/database"
	"user-service/src/interface/repository"
	"user-service/src/mock/cache"
	"user-service/src/model/dto"
	repoimpl "user-service/src/repository"
	"user-service/src/test/util"

	"github.com/jackc/pgx/v5/pgxpool"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/repository/test/... -count=1
// go test -run ^TestRepository_CreateUser$ -v ./src/repository/test/ -count=1

type CreateUserTestSuite struct {
	suite.Suite
	userRepo     repository.User
	userCache    *cache.UserMock
	postgresDB   *pgxpool.Pool
	userTestUtil *util.UserTest
}

func (t *CreateUserTestSuite) SetupSuite() {
	t.postgresDB = database.NewPostgres()
	t.userCache = cache.NewUserMock()

	t.userTestUtil = util.NewUserTest(t.postgresDB)
	t.userRepo = repoimpl.NewUser(t.postgresDB, t.userCache)
}

func (t *CreateUserTestSuite) TearDownTest() {
	t.userTestUtil.Delete()
}

func (t *CreateUserTestSuite) TearDownSuite() {
	t.postgresDB.Close()
}

func (t *CreateUserTestSuite) Test_Success() {
	id, err := gonanoid.New()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "test.CreateUserTestSuite/Test_Success", "section": "gonanoid.New"}).Error(err)
		return
	}
	err = t.userRepo.Create(context.Background(), &dto.RegisterReq{
		UserId:   id,
		Email:    "johndoe123@gmail.com",
		FullName: "John Doe",
		Password: "$2a$10$Rx3JG/oPRbY1urgS8CgqauFkGEeAcE15IKBCSuDXKqsdDy8GZlz1y",
	})

	assert.NoError(t.T(), err)
}

func TestRepository_CreateUser(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
