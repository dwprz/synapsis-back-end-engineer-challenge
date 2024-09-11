package cache

import (
	"user-service/src/model/entity"
	"context"
)

type User interface {
	Cache(ctx context.Context, user *entity.User)
	FindByEmail(ctx context.Context, email string) *entity.User
	DeleteByEmail(ctx context.Context, email string)
}
