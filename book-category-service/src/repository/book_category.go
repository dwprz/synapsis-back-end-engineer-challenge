package repository

import (
	"book-category-service/src/common/errors"
	"book-category-service/src/common/helper"
	"book-category-service/src/interface/repository"
	"book-category-service/src/model/dto"
	"book-category-service/src/model/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookCategoryImpl struct {
	pool *pgxpool.Pool
}

func NewBookCategory(p *pgxpool.Pool) repository.BookCategory {
	return &BookCategoryImpl{
		pool: p,
	}
}

func (r *BookCategoryImpl) Create(ctx context.Context, data *dto.CreateBookCategoryReq) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query, args, err := helper.BuildCreateBookCategoryQuery(data)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, args...)
	return err
}

func (r *BookCategoryImpl) FindManyByCategory(ctx context.Context, category entity.Category, limit, offset int) (bookIds []int, err error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := `SELECT book_id FROM book_categories WHERE category = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;`

	rows, err := conn.Query(ctx, query, category, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var bookId int
		err := rows.Scan(&bookId)
		if err != nil {
			return nil, err
		}

		bookIds = append(bookIds, bookId)
	}

	if len(bookIds) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "no books match this category"}
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return bookIds, nil
}

func (r *BookCategoryImpl) Delete(ctx context.Context, data *dto.DeleteBookCategoryReq) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query, args, err := helper.BuildDeleteBookCategoryQuery(data)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, args...)
	return err
}
