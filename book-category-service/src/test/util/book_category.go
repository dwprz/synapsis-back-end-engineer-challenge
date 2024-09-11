package util

import (
	"book-category-service/src/common/log"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type BookCategoryTest struct {
	pool *pgxpool.Pool
}

func NewBookCategoryTest(p *pgxpool.Pool) *BookCategoryTest {
	return &BookCategoryTest{
		pool: p,
	}
}

func (u *BookCategoryTest) Delete() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookCategoryTest/Delete", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	if _, err = conn.Exec(ctx, `DELETE FROM book_categories`); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookCategoryTest/Delete", "section": "conn.Exec"})
	}
}

func (u *BookCategoryTest) CreateMany() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookCategoryTest/CreateMany", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	query := `
	INSERT INTO
		book_categories(category, book_id)
	VALUES
		('ADVENTURE', 1), ('ADVENTURE', 2), ('FICTION', 3), ('HISTORY', 4), ('ADVENTURE', 5);
	`

	if _, err = conn.Exec(ctx, query); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookCategoryTest/CreateMany", "section": "conn.Exec"})
	}
}
