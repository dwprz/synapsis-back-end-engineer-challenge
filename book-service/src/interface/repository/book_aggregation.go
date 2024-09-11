package repository

import (
	"book-service/src/model/dto"
	"context"
)

type BookAggregation interface {
	CountByTitle(ctx context.Context, title string) (totalBooks int, err error)
	CountByFields(ctx context.Context, data *dto.GetBookReq) (totalBooks int, err error)
	CountAll(ctx context.Context) (totalBooks int, err error)
}
