package repository

import (
	"book-service/src/model/entity"
	"context"
)

type PopularTitleKey interface {
	Upsert(ctx context.Context, title_key string)
	FindManyPopularBook(ctx context.Context) ([]*entity.Book, error)
}
