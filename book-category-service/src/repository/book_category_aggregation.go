package repository

import (
	"book-category-service/src/interface/repository"
	"book-category-service/src/model/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookCategoryAggregationImpl struct {
	pool *pgxpool.Pool
}

func NewBookCategoryAggregation(p *pgxpool.Pool) repository.BookCategoryAggregation {
	return &BookCategoryAggregationImpl{
		pool: p,
	}
}

func (r *BookCategoryAggregationImpl) CountByCategory(ctx context.Context, category entity.Category) (totalBooks int, err error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	query := `SELECT COUNT(*) FROM book_categories WHERE category = $1;`
	err = conn.QueryRow(ctx, query, category).Scan(&totalBooks)

	return totalBooks, err
}
