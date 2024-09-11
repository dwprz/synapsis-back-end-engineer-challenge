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
// go test -run ^TestService_UpdateUser$ -v ./src/service/test/ -count=1

type UpdateUserTestSuite struct {
	suite.Suite
	userService service.User
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *UpdateUserTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.userService = serviceimpl.NewUser(t.userRepo, t.userCache)
}

func (t *UpdateUserTestSuite) Test_Succcess() {
	email := "johndoe123@gmail.com"
	req := t.createUpdateUserReq()
	req.Email = email

	user := t.createEntityUser("rahasia")

	t.MockUserCache_FindByEmail(email, nil)
	t.MockUserRepo_FindByFields(&entity.User{Email: email}, nil, nil)

	t.MockUserRepo_UpdateByUserId(user.UserId, user, nil)

	res, err := t.userService.Update(context.Background(), req)
	assert.NoError(t.T(), err)

	assert.Equal(t.T(), email, res.Email)
}

func (t *UpdateUserTestSuite) MockUserCache_FindByEmail(email string, returnArg1 *entity.User) {

	t.userCache.Mock.On("FindByEmail", mock.Anything, email).Return(returnArg1)
}

func (t *UpdateUserTestSuite) MockUserRepo_FindByFields(user *entity.User, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("FindByFields", mock.Anything, user).Return(returnArg1, returnArg2)
}

func (t *UpdateUserTestSuite) MockUserRepo_UpdateByUserId(userId string, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("UpdateByUserId", mock.Anything, mock.MatchedBy(func(args *dto.UpdateUserReq) bool {
		return userId == args.UserId && args.Email != ""
	})).Return(returnArg1, returnArg2)
}

func (t *UpdateUserTestSuite) createUpdateUserReq() *dto.UpdateUserReq {
	return &dto.UpdateUserReq{
		UserId: "ynA1nZIULkXLrfy0fvz5t",
	}
}

func (t *UpdateUserTestSuite) createEntityUser(password string) *entity.User {
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
func TestService_UpdateUser(t *testing.T) {
	suite.Run(t, new(UpdateUserTestSuite))
}
