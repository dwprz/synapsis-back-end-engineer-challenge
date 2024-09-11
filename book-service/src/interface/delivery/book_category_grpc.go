package delivery

import "context"

type BookCategoryGrpc interface {
	DeleteBookFromCategoryReq(ctx context.Context, bookId int) error
}


