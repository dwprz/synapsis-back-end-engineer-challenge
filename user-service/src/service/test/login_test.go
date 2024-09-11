package test

import (
	"context"
	"testing"
	"time"
	"user-service/src/common/log"
	"user-service/src/interface/service"
	"user-service/src/mock/cache"
	"user-service/src/mock/repository"
	"user-service/src/model/dto"
	"user-service/src/model/entity"
	serviceimpl "user-service/src/service"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_Login$ -v ./src/service/test/ -count=1

type LoginTestSuite struct {
	suite.Suite
	authService service.Auth
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *LoginTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.authService = serviceimpl.NewAuth(t.userRepo, t.userCache)
}

func (t *LoginTestSuite) Test_Succcess() {
	req := t.createLoginReq()

	user := t.createEntityUser(req.Password)

	t.MockUserCache_FindByEmail(req.Email, user)
	t.MockUserRepo_UpdateByUserId(user.UserId, nil, nil)

	res, err := t.authService.Login(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), res.Data.Email, req.Email)
	assert.NotEmpty(t.T(), res.Tokens.AccessToken)
	assert.NotEmpty(t.T(), res.Tokens.RefreshToken)
}

func (t *LoginTestSuite) MockUserCache_FindByEmail(email string, returnArg1 *entity.User) {

	t.userCache.Mock.On("FindByEmail", mock.Anything, email).Return(returnArg1)
}

func (t *LoginTestSuite) MockUserRepo_FindByFields(user *entity.User, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("FindByFields", mock.Anything, user).Return(returnArg1, returnArg2)
}

func (t *LoginTestSuite) MockUserRepo_UpdateByUserId(userId string, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("UpdateByUserId", mock.Anything, mock.MatchedBy(func (args *dto.UpdateUserReq) bool  {
		return userId == args.UserId && args.RefreshToken != ""
	})).Return(returnArg1, returnArg2)
}

func (t *LoginTestSuite) createLoginReq() *dto.LoginReq {
	return &dto.LoginReq{
		Email:    "johndoe123@gmail.com",
		Password: "rahasia",
	}
}

func (t *LoginTestSuite) createEntityUser(password string) *entity.User {
	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "database.NewPostgres", "section": "pgxpool.NewWithConfig"}).Fatal(err)
	}

	return &entity.User{
		UserId:    "ynA1nZIULkXLrfy0fvz5t",
		Email:     "johndoe123@gmail.com",
		FullName:  "John Doe",
		Role:      "user",
		Password:  string(encryptPwd),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func TestService_Login(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
