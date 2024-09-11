package util

import (
	"book-service/src/common/log"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type BookTest struct {
	pool *pgxpool.Pool
}

func NewBookTest(p *pgxpool.Pool) *BookTest {
	return &BookTest{
		pool: p,
	}
}

func (u *BookTest) Delete() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/Delete", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	if _, err = conn.Exec(ctx, `DELETE FROM books`); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/Delete", "section": "conn.Exec"})
	}

	if _, err = conn.Exec(ctx, `DELETE FROM popular_title_keys`); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/Delete", "section": "conn.Exec"})
	}
}

func (u *BookTest) CreateMany() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/CreateMany", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	query := `
	INSERT INTO books (book_id, title, author, isbn, synopsis, published_year, stock, location)
	VALUES
	    (1, 'Amazing Adventure', 'Arthur Conan Doyle', '978-0-123456-47-2', 'A thrilling adventure novel.', 2024, 10, 'Shelf A'),
	    (2, 'Fictional Story', 'J.K. Rowling', '978-0-123456-48-9', 'A captivating fictional story.', 2023, 15, 'Shelf B'),
	    (3, 'History', 'Stephen Hawking', '978-0-123456-49-6', 'An insightful history book.', 2022, 5, 'Shelf C'),
	    (4, 'Abiut Science', 'Carl Sagan', '978-0-123456-50-3', 'A book on scientific phenomena.', 2021, 8, 'Shelf D'),
	    (5, 'Fantasy III', 'J.R.R. Tolkien', '978-0-123456-51-0', 'A magical fantasy adventure.', 2020, 12, 'Shelf E');
	`

	if _, err = conn.Exec(ctx, query); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/CreateMany", "section": "conn.Exec"})
		return
	}
}

func (u *BookTest) AddPopularTitleKey() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/AddPopularTitleKey", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	query := `
	INSERT INTO 
		popular_title_keys(title_key, total_search) 
	VALUES 
		('amazing adventure', 1000);
	`

	if _, err = conn.Exec(ctx, query); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.BookTest/AddPopularTitleKey", "section": "conn.Exec"})
		return
	}
}

