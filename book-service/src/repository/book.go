package repository

import (
	"book-service/src/api/grpc/client"
	errcstm "book-service/src/common/errors"
	"book-service/src/common/helper"
	"book-service/src/interface/repository"
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	"context"
	"errors"
	"strings"
	"time"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookImpl struct {
	pool       *pgxpool.Pool
	grpcClient *client.Grpc
}

func NewBook(p *pgxpool.Pool, gc *client.Grpc) repository.Book {
	return &BookImpl{
		pool:       p,
		grpcClient: gc,
	}
}

func (r *BookImpl) Create(ctx context.Context, data *dto.AddBookReq) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := `
	INSERT INTO 
		books(title, author, isbn, synopsis, published_year, stock, location)
	VALUES
		($1, $2, $3, $4, $5, $6, $7);
	`

	_, err = conn.Exec(ctx, query, data.Title, data.Author, data.ISBN, data.Synopsis, data.PublishedYear, data.Stock, data.Location)
	return err
}

func (r *BookImpl) FindByTitle(ctx context.Context, title string) (*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := `
	SELECT  
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at 
	FROM 
		books WHERE title = $1;
	`

	row := conn.QueryRow(ctx, query, title)

	var book entity.Book
	row.Scan(&book.BookId, &book.Title, &book.Author, &book.ISBN, &book.Synopsis, &book.PublishedYear, &book.Stock,
		&book.Location, &book.CreatedAt, &book.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &book, err
}

func (r *BookImpl) FindManyByTitle(ctx context.Context, title string, limit, offset int) ([]*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	title = strings.Join(strings.Split(title, " "), " & ")

	query := `
	SELECT
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at 
	FROM
		books
	WHERE
		to_tsvector('indonesian', title) @@ to_tsquery('indonesian', $1)
    LIMIT 
		$2 OFFSET $3;
	`

	rows, err := conn.Query(ctx, query, title, limit, offset)
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
		return nil, &errcstm.Response{HttpCode: 404, Message: "books not found"}
	}

	return books, err
}

func (r *BookImpl) FindManyByFields(ctx context.Context, data *dto.GetBookReq, limit, offset int) ([]*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query, args, err := helper.BuildFindManyByFieldsQuery(data, limit, offset)
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, query, args...)
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
		return nil, &errcstm.Response{HttpCode: 404, Message: "books not found"}
	}

	return books, err
}

func (r *BookImpl) FindManyByRandom(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := `
	SELECT 
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at 
	FROM 
		books
	ORDER 
		BY RANDOM()
	LIMIT $1 OFFSET $2;
	`

	rows, err := conn.Query(ctx, query, limit, offset)
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
		return nil, &errcstm.Response{HttpCode: 404, Message: "books not found"}
	}

	return books, err
}

func (r *BookImpl) FindManyByIds(ctx context.Context, bookIds []uint32) ([]*pb.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query, args, err := helper.BuildFindManyByIdsQuery(bookIds)
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*pb.Book
	for rows.Next() {
		var (
			book      pb.Book
			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.Isbn, &book.Synopsis, &book.PublishedYear, &book.Stock,
			&book.Location, &createdAt, &updatedAt)

		if err != nil {
			return nil, err
		}

		book.CreatedAt = timestamppb.New(createdAt)
		book.UpdatedAt = timestamppb.New(updatedAt)

		books = append(books, &book)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, &errcstm.Response{HttpCode: 404, Message: "books not found"}
	}

	return books, err
}

func (r *BookImpl) UpdateById(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	if data.Stock != nil {
		_, err := tx.Exec(ctx, `SELECT stock FROM books WHERE book_id = $1 FOR UPDATE;`, data.BookId)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
	}

	query, args, err := helper.BuildUpdateByIdQuery(data)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(ctx, query, args...)

	var book entity.Book
	err = row.Scan(&book.BookId, &book.Title, &book.Author, &book.ISBN, &book.Synopsis, &book.PublishedYear, &book.Stock,
		&book.Location, &book.CreatedAt, &book.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		tx.Rollback(ctx)
		return nil, &errcstm.Response{HttpCode: 404, Message: "book not found"}
	}

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	tx.Commit(ctx)

	return &book, nil
}

func (r *BookImpl) DeleteById(ctx context.Context, bookId int) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	query := `DELETE FROM books WHERE book_id = $1;`
	if _, err = conn.Exec(ctx, query, bookId); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := r.grpcClient.BookCategory.DeleteBookFromCategoryReq(ctx, bookId); err != nil {
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)

	return err
}
