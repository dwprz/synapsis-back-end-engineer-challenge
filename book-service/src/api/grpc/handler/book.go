package handler

import (
	"book-service/src/interface/service"
	"book-service/src/model/dto"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type BookGrpcImpl struct {
	bookService service.Book
	pb.UnimplementedBookServiceServer
}

func NewBookGrpc(bs service.Book) pb.BookServiceServer {
	return &BookGrpcImpl{
		bookService: bs,
	}
}

func (h *BookGrpcImpl) FindManyByIds(ctx context.Context, data *pb.BookIds) (*pb.FindManyRes, error) {
	books, err := h.bookService.FindManyByIds(ctx, &dto.FindManyByIdsReq{BookIds: data.Ids})

	return &pb.FindManyRes{Data: books}, err
}