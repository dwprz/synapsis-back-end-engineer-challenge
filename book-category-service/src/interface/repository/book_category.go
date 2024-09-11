package repository

import (
	"book-category-service/src/model/dto"
	"book-category-service/src/model/entity"
	"context"
)

type BookCategory interface {
	Create(ctx context.Context, data *dto.CreateBookCategoryReq) error
	FindManyByCategory(ctx context.Context, category entity.Category, limit, offset int) (bookIds []int, err error)
	Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error
}
