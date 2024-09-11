package service

import (
	"context"
	"user-service/src/model/dto"
	"user-service/src/model/entity"
)

type User interface {
	FindByEmail(ctx context.Context, email string) (*entity.SanitizedUser, error)
	Update(ctx context.Context, data *dto.UpdateUserReq) (*entity.SanitizedUser, error)
	Delete(ctx context.Context, userId string) error  
}
