package service

import (
	"context"
	"user-service/src/model/dto"
	"user-service/src/model/entity"
)

type Auth interface {
	Register(ctx context.Context, data *dto.RegisterReq) error
	Login(ctx context.Context, data *dto.LoginReq) (*dto.LoginRes, error)
	RefreshToken(ctx context.Context, refreshToken string) (*entity.Tokens, error)
	SetNullRefreshToken(ctx context.Context, refreshToken string) error
}
