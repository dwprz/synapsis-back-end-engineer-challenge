package repository

import (
	"context"
	"user-service/src/model/dto"
	"user-service/src/model/entity"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func NewUserMock() *UserMock {
	return &UserMock{
		Mock: mock.Mock{},
	}
}

func (u *UserMock) Create(ctx context.Context, data *dto.RegisterReq) error {
	arguments := u.Mock.Called(ctx, data)

	return arguments.Error(0)
}

func (u *UserMock) FindByFields(ctx context.Context, data *entity.User) (*entity.User, error) {
	arguments := u.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.User), arguments.Error(1)
}

func (u *UserMock) UpdateByUserId(ctx context.Context, data *dto.UpdateUserReq) (*entity.User, error) {
	arguments := u.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.User), arguments.Error(1)
}

func (u *UserMock) SetNullRefreshToken(ctx context.Context, refreshToken string) error {
	arguments := u.Mock.Called(ctx, refreshToken)

	return arguments.Error(0)
}

func (u *UserMock) Delete(ctx context.Context, userId string) error {
	arguments := u.Mock.Called(ctx, userId)

	return arguments.Error(0)
}
