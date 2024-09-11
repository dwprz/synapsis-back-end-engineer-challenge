package test

import (
	"context"
	"testing"
	"user-service/src/common/errors"
	"user-service/src/interface/service"
	"user-service/src/mock/cache"
	"user-service/src/mock/repository"
	"user-service/src/model/dto"
	"user-service/src/model/entity"
	serviceimpl "user-service/src/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_Register$ -v ./src/service/test/ -count=1

type RegisterTestSuite struct {
	suite.Suite
	authService service.Auth
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *RegisterTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.authService = serviceimpl.NewAuth(t.userRepo, t.userCache)
}

func (t *RegisterTestSuite) Test_Succcess() {
	req := t.createRegisterReq()

	t.MockUserCache_FindByEmail(req.Email, nil)
	t.MockUserRepo_FindByFields(&entity.User{Email: req.Email}, nil, nil)

	t.MockUserRepo_Create(req, nil)

	err := t.authService.Register(context.Background(), req)
	assert.NoError(t.T(), err)
}

func (t *RegisterTestSuite) Test_AlreadyExists() {
	req := t.createRegisterReq()
	req.Email = "useralreadyexist@gmail.com"

	t.MockUserCache_FindByEmail(req.Email, &entity.User{
		FullName: req.FullName,
		Email:    req.Email,
	})

	err := t.authService.Register(context.Background(), req)
	assert.Error(t.T(), err)

	errRes, ok := err.(*errors.Response)
	assert.True(t.T(), ok)
	assert.Equal(t.T(), 409, errRes.HttpCode)
}

func (t *RegisterTestSuite) MockUserCache_FindByEmail(email string, returnArg1 *entity.User) {

	t.userCache.Mock.On("FindByEmail", mock.Anything, email).Return(returnArg1)
}

func (t *RegisterTestSuite) MockUserRepo_FindByFields(user *entity.User, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("FindByFields", mock.Anything, user).Return(returnArg1, returnArg2)
}

func (t *RegisterTestSuite) MockUserRepo_Create(data *dto.RegisterReq, returnArg1 error) {

	t.userRepo.Mock.On("Create", mock.Anything, mock.MatchedBy(func(args *dto.RegisterReq) bool {
		return data.Email == args.Email && data.FullName == args.FullName && args.Password != ""

	})).Return(returnArg1)
}

func (t *RegisterTestSuite) createRegisterReq() *dto.RegisterReq {
	return &dto.RegisterReq{
		Email:    "johndoe123@gmail.com",
		FullName: "John Doe",
		Password: "rahasia",
	}
}

func TestService_Register(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
