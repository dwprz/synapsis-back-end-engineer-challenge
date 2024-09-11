package repository

import (
	"book-category-service/src/model/entity"
	"context"
)

type BookCategoryAggregation interface {
	CountByCategory(ctx context.Context, category entity.Category) (totalBooks int, err error)
}
