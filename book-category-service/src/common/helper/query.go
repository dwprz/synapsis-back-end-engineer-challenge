package helper

import (
	"book-category-service/src/common/errors"
	"book-category-service/src/model/dto"
	"fmt"
	"strings"
)

func BuildCreateBookCategoryQuery(data *dto.CreateBookCategoryReq) (query string, args []any, err error) {
	if len(data.BookIds) == 0 {
		return "", nil, &errors.Response{HttpCode: 400, Message: "book ids are missing, cannot create book category query"}
	}

	var placeholders []string
	argId := 1 

	for _, id := range data.BookIds {
		placeholder := fmt.Sprintf("($%d, $%d)", argId, argId+1)
		placeholders = append(placeholders, placeholder)
		
		args = append(args, data.Category, id)
		argId += 2
	}

	query = fmt.Sprintf(`
	INSERT INTO 
		book_categories(category, book_id) 
	VALUES 
		%s 
	ON CONFLICT 
        (category, book_id) 
    DO UPDATE SET 
        category = EXCLUDED.category, book_id = EXCLUDED.book_id;`,
		strings.Join(placeholders, ", "))

	return query, args, nil
}

func BuildDeleteBookCategoryQuery(data *dto.DeleteBookCategoryReq) (query string, args []any, err error) {
	if data == nil {
		return "", nil, &errors.Response{HttpCode: 400, Message: "fields are missing, cannot create delete book category query"}
	}

	var whereClauses []string
	argId := 1

	if data.Category != nil {
		clause := fmt.Sprintf(`category = $%d`, argId)
		whereClauses = append(whereClauses, clause)
		args = append(args, *data.Category)
		argId++
	}

	if data.BookId != nil {
		clause := fmt.Sprintf(`book_id = $%d`, argId)
		whereClauses = append(whereClauses, clause)
		args = append(args, *data.BookId)
		argId++
	}

	query = fmt.Sprintf("DELETE FROM book_categories WHERE %s;", strings.Join(whereClauses, " AND "))

	return query, args, nil
}
