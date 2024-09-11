package service

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

func (s *UserMock) FindByEmail(ctx context.Context, email string) (*entity.SanitizedUser, error) {
	arguments := s.Mock.Called(ctx, email)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.SanitizedUser), arguments.Error(1)
}

func (u *UserMock) Update(ctx context.Context, data *dto.UpdateUserReq) (*entity.SanitizedUser, error) {
	arguments := u.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.SanitizedUser), arguments.Error(1)
}


func (u *UserMock) Delete(ctx context.Context, userId string) error {
	arguments := u.Mock.Called(ctx, userId)

	return arguments.Error(0)
}
