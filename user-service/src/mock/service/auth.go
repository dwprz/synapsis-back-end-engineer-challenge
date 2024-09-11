package service

import (
	"context"
	"user-service/src/model/dto"
	"user-service/src/model/entity"

	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (s *AuthMock) Register(ctx context.Context, data *dto.RegisterReq) error {
	arguments := s.Mock.Called(ctx, data)

	return arguments.Error(0)
}

func (s *AuthMock) Login(ctx context.Context, data *dto.LoginReq) (*dto.LoginRes, error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.LoginRes), arguments.Error(1)
}

func (s *AuthMock) RefreshToken(ctx context.Context, refreshToken string) (*entity.Tokens, error) {
	arguments := s.Mock.Called(ctx, refreshToken)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.Tokens), arguments.Error(1)
}

func (s *AuthMock) SetNullRefreshToken(ctx context.Context, refreshToken string) error {
	arguments := s.Mock.Called(ctx, refreshToken)

	return arguments.Error(0)
}
