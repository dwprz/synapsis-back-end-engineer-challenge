package repository

import (
	"book-service/src/common/helper"
	"book-service/src/interface/repository"
	"book-service/src/model/dto"
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookAggregationImpl struct {
	pool *pgxpool.Pool
}

func NewBookAggregation(p *pgxpool.Pool) repository.BookAggregation {
	return &BookAggregationImpl{
		pool: p,
	}
}

func (r *BookAggregationImpl) CountByTitle(ctx context.Context, title string) (totalBooks int, err error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	title = strings.Join(strings.Split(title, " "), " & ")

	query := `
	SELECT
		COUNT(*) 
	FROM
		books
	WHERE
		to_tsvector('indonesian', title) @@ to_tsquery('indonesian', $1);
	`

	err = conn.QueryRow(ctx, query, title).Scan(&totalBooks)
	return totalBooks, err
}

func (r *BookAggregationImpl) CountByFields(ctx context.Context, data *dto.GetBookReq) (totalBooks int, err error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	query, args, err := helper.BuildCountByFieldsQuery(data)
	if err != nil {
		return 0, err
	}

	err = conn.QueryRow(ctx, query, args...).Scan(&totalBooks)
	return totalBooks, err
}

func (r *BookAggregationImpl) CountAll(ctx context.Context) (totalBooks int, err error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	query := `SELECT COUNT(*) FROM books;`
	err = conn.QueryRow(ctx, query).Scan(&totalBooks)

	return totalBooks, err
}
