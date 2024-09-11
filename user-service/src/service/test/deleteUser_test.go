package test

import (
	"context"
	"testing"
	"user-service/src/interface/service"
	"user-service/src/mock/cache"
	"user-service/src/mock/repository"
	"user-service/src/model/entity"
	serviceimpl "user-service/src/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestService_DeleteUser$ -v ./src/service/test/ -count=1

type DeleteUserTestSuite struct {
	suite.Suite
	userService service.User
	userRepo    *repository.UserMock
	userCache   *cache.UserMock
}

func (t *DeleteUserTestSuite) SetupSuite() {
	t.userRepo = repository.NewUserMock()
	t.userCache = cache.NewUserMock()

	t.userService = serviceimpl.NewUser(t.userRepo, t.userCache)
}

func (t *DeleteUserTestSuite) Test_Succcess() {
	userId := "ynA1nZIULkXLrfy0fvz5t"

	t.MockUserRepo_Delete(userId, nil)

	err := t.userService.Delete(context.Background(), userId)
	assert.NoError(t.T(), err)
}

func (t *DeleteUserTestSuite) MockUserCache_FindByEmail(email string, returnArg1 *entity.User) {

	t.userCache.Mock.On("FindByEmail", mock.Anything, email).Return(returnArg1)
}

func (t *DeleteUserTestSuite) MockUserRepo_FindByFields(user *entity.User, returnArg1 *entity.User, returnArg2 error) {

	t.userRepo.Mock.On("FindByFields", mock.Anything, user).Return(returnArg1, returnArg2)
}

func (t *DeleteUserTestSuite) MockUserRepo_Delete(userId string, returnArg1 error) {

	t.userRepo.Mock.On("Delete", mock.Anything, userId).Return(returnArg1)
}

func TestService_DeleteUser(t *testing.T) {
	suite.Run(t, new(DeleteUserTestSuite))
}
