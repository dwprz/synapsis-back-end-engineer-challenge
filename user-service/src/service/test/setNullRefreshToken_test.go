package test

import (
	"context"
	"testing"
	"user-service/src/interface/service"
	"user-service/src/mock/cache"
	"user-service/src/mock/repository"
	serviceimpl "user-service/src/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_SetNullRefreshToken$ -v ./src/service/test/ -count=1

type SetNullRefreshTokenTestSuite struct {
	suite.Suite
	authService service.Auth
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *SetNullRefreshTokenTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.authService = serviceimpl.NewAuth(t.userRepo, t.userCache)
}

func (t *SetNullRefreshTokenTestSuite) Test_Succcess() {
	refreshToken := "example-refresh-token"

	t.MockUserRepo_SetNullRefreshToken(refreshToken, nil)

	err := t.authService.SetNullRefreshToken(context.Background(), refreshToken)
	assert.NoError(t.T(), err)
}

func (t *SetNullRefreshTokenTestSuite) MockUserRepo_SetNullRefreshToken(refreshToken string, returnArg1 error) {

	t.userRepo.Mock.On("SetNullRefreshToken", mock.Anything, refreshToken).Return(returnArg1)
}

func TestService_SetNullRefreshToken(t *testing.T) {
	suite.Run(t, new(SetNullRefreshTokenTestSuite))
}
