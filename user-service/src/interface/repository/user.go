package repository

import (
	"user-service/src/model/dto"
	"user-service/src/model/entity"
	"context"
)

type User interface {
	Create(ctx context.Context, data *dto.RegisterReq) error
	FindByFields(ctx context.Context, data *entity.User) (*entity.User, error)
	UpdateByUserId(ctx context.Context, data *dto.UpdateUserReq) (*entity.User, error)
	SetNullRefreshToken(ctx context.Context, refreshToken string) error
	Delete(ctx context.Context, userId string) error
}
