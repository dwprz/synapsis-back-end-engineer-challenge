package delivery

import (
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type BookGrpc interface {
	FindManyByIds(ctx context.Context, bookIds []int) ([]*pb.Book, error)
}
