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
// go test -run ^TestService_RefreshToken$ -v ./src/service/test/ -count=1

type RefreshTokenTestSuite struct {
	suite.Suite
	authService service.Auth
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *RefreshTokenTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.authService = serviceimpl.NewAuth(t.userRepo, t.userCache)
}

func (t *RefreshTokenTestSuite) Test_Succcess() {
	refreshToken := `example-refresh-token`

	user := t.createEntityUser("rahasia")

	t.MockUserRepo_FindByFields(&entity.User{RefreshToken: &refreshToken}, user, nil)
	t.MockUserRepo_UpdateByUserId(user.UserId, mock.Anything, nil, nil)

	res, err := t.authService.RefreshToken(context.Background(), refreshToken)
	assert.NoError(t.T(), err)

	assert.NotEmpty(t.T(), res.AccessToken)
}

func (t *RefreshTokenTestSuite) MockUserCache_FindByEmail(email string, returnArg1 *entity.User) {

	t.userCache.Mock.On("FindByEmail", mock.Anything, email).Return(returnArg1)
}

func (t *RefreshTokenTestSuite) MockUserRepo_FindByFields(user *entity.User, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("FindByFields", mock.Anything, user).Return(returnArg1, returnArg2)
}

func (t *RefreshTokenTestSuite) MockUserRepo_UpdateByUserId(userId string, refreshToken string, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("UpdateByUserId", mock.Anything, &dto.UpdateUserReq{UserId: userId, RefreshToken: refreshToken}).Return(returnArg1, returnArg2)
}

func (t *RefreshTokenTestSuite) createEntityUser(password string) *entity.User {

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
func TestService_RefreshToken(t *testing.T) {
	suite.Run(t, new(RefreshTokenTestSuite))
}
