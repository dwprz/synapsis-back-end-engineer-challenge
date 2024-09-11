package helper

import (
	"book-service/src/common/errors"
	"book-service/src/model/dto"
	"fmt"
	"strings"
)

func BuildFindManyByFieldsQuery(data *dto.GetBookReq, limit, offset int) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	var whereClause []string
	argId := 1

	if data.BookId != 0 {
		whereClause = append(whereClause, fmt.Sprintf("book_id = $%d", argId))
		args = append(args, data.BookId)
		argId++
	}

	if data.Author != "" {
		whereClause = append(whereClause, fmt.Sprintf("author = $%d", argId))
		args = append(args, data.Author)
		argId++
	}

	if data.ISBN != "" {
		whereClause = append(whereClause, fmt.Sprintf("isbn = $%d", argId))
		args = append(args, data.ISBN)
		argId++
	}

	if data.PublishedYear != 0 {
		whereClause = append(whereClause, fmt.Sprintf("published_year = $%d", argId))
		args = append(args, data.PublishedYear)
		argId++
	}

	if data.Stock != nil {
		whereClause = append(whereClause, fmt.Sprintf("stock = $%d", argId))
		args = append(args, *data.Stock)
		argId++
	}

	if len(whereClause) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	query = fmt.Sprintf(`
	SELECT 
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at 
	FROM 
		books WHERE %s LIMIT $%d OFFSET $%d;
	`, strings.Join(whereClause, " AND "), argId, argId+1)

	args = append(args, limit, offset)

	return query, args, nil
}

func BuildCountByFieldsQuery(data *dto.GetBookReq) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	var whereClause []string
	argId := 1

	if data.BookId != 0 {
		whereClause = append(whereClause, fmt.Sprintf("book_id = $%d", argId))
		args = append(args, data.BookId)
		argId++
	}

	if data.Author != "" {
		whereClause = append(whereClause, fmt.Sprintf("author = $%d", argId))
		args = append(args, data.Author)
		argId++
	}

	if data.ISBN != "" {
		whereClause = append(whereClause, fmt.Sprintf("isbn = $%d", argId))
		args = append(args, data.ISBN)
		argId++
	}

	if data.PublishedYear != 0 {
		whereClause = append(whereClause, fmt.Sprintf("published_year = $%d", argId))
		args = append(args, data.PublishedYear)
		argId++
	}

	if data.Stock != nil {
		whereClause = append(whereClause, fmt.Sprintf("stock = $%d", argId))
		args = append(args, *data.Stock)
		argId++
	}

	if len(whereClause) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	query = fmt.Sprintf(`SELECT COUNT(*) FROM books WHERE %s;`, strings.Join(whereClause, " AND "))

	return query, args, nil
}

func BuildUpdateByIdQuery(data *dto.UpdateBookReq) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	var whereClause []string
	argId := 1

	if data.Title != "" {
		whereClause = append(whereClause, fmt.Sprintf("title = $%d", argId))
		args = append(args, data.Title)
		argId++
	}

	if data.Author != "" {
		whereClause = append(whereClause, fmt.Sprintf("author = $%d", argId))
		args = append(args, data.Author)
		argId++
	}

	if data.ISBN != "" {
		whereClause = append(whereClause, fmt.Sprintf("isbn = $%d", argId))
		args = append(args, data.ISBN)
		argId++
	}

	if data.Synopsis != nil {
		whereClause = append(whereClause, fmt.Sprintf("synopsis = $%d", argId))
		args = append(args, *data.Synopsis)
		argId++
	}

	if data.PublishedYear != 0 {
		whereClause = append(whereClause, fmt.Sprintf("published_year = $%d", argId))
		args = append(args, data.PublishedYear)
		argId++
	}

	if data.Stock != nil {
		whereClause = append(whereClause, fmt.Sprintf("stock = $%d", argId))
		args = append(args, *data.Stock)
		argId++
	}

	if data.Location != "" {
		whereClause = append(whereClause, fmt.Sprintf("location = $%d", argId))
		args = append(args, data.Location)
		argId++
	}

	if len(whereClause) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no fields to select"}
	}

	query = fmt.Sprintf(`
	UPDATE 
		books SET %s 
	WHERE 
		book_id = $%d 
	RETURNING 
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at;
	`, strings.Join(whereClause, ", "), argId)

	args = append(args, data.BookId)

	return query, args, nil
}

func BuildFindManyByIdsQuery(bookIds []uint32) (query string, args []any, err error) {
	if len(bookIds) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "no book book ids provided"}
	}

	var placeholders []string
	argId := 1
	
	for _, id := range bookIds {
		placeholder := fmt.Sprintf("$%d", argId)
		placeholders = append(placeholders, placeholder)
		args = append(args, id)
		argId++
	}

	query = fmt.Sprintf(`
	SELECT 
		book_id, title, author, isbn, synopsis, published_year, stock, location, created_at, updated_at 
	FROM 
		books 
	WHERE 
		book_id IN (%s); 
	`, strings.Join(placeholders, ", "))

	return query, args, nil
}
