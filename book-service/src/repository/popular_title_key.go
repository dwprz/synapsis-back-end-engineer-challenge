package repository

import (
	"book-service/src/common/errors"
	"book-service/src/common/log"
	"book-service/src/interface/repository"
	"book-service/src/model/entity"
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PopularTitleKeyImpl struct {
	pool *pgxpool.Pool
}

func NewPopularTitleKey(p *pgxpool.Pool) repository.PopularTitleKey {
	return &PopularTitleKeyImpl{
		pool: p,
	}
}

func (r *PopularTitleKeyImpl) Upsert(ctx context.Context, title_key string) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "repository.PopularTitleKeyImpl/Upsert", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	title_key = strings.ToLower(title_key)

	query := `
	INSERT INTO 
		popular_title_keys(title_key) 
	VALUES 
		($1) 
	ON CONFLICT 
        (title_key) 
    DO UPDATE SET 
		 total_search = EXCLUDED.total_search + 1;
	`

	if _, err = conn.Exec(ctx, query, title_key); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "repository.PopularTitleKeyImpl/Upsert", "section": "conn.Exec"}).Error(err)
	}
}

func (r *PopularTitleKeyImpl) FindManyPopularBook(ctx context.Context) ([]*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := `
	WITH cte_popular_title_key AS (
	    SELECT title_key FROM popular_title_keys ORDER BY total_search DESC LIMIT 1
	)
	SELECT 
	    book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at
	FROM
	    books
	WHERE
	    to_tsvector('indonesian', title) @@ to_tsquery(
			'indonesian', 
			replace((SELECT title_key FROM cte_popular_title_key), ' ', ' & ')
		)
	LIMIT 20 OFFSET 0;
	`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.ISBN, &book.Synopsis, &book.PublishedYear, &book.Stock,
			&book.Location, &book.CreatedAt, &book.UpdatedAt)

		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "popular books not found"}
	}

	return books, err
}
