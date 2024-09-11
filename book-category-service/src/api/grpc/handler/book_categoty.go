package handler

import (
	"book-category-service/src/interface/service"
	"book-category-service/src/model/dto"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book_category"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookCategoryGrpcImpl struct {
	bookCategoryService service.BookCategory
	pb.UnimplementedBookCategoryServiceServer
}

func NewBookCategoryGrpc(bs service.BookCategory) pb.BookCategoryServiceServer {
	return &BookCategoryGrpcImpl{
		bookCategoryService: bs,
	}
}

func (h *BookCategoryGrpcImpl) DeleteBookFromCategoryReq(ctx context.Context, data *pb.BookId) (*emptypb.Empty, error) {
	bookId := int(data.Id)
	err := h.bookCategoryService.Delete(ctx, &dto.DeleteBookCategoryReq{BookId: &bookId})

	return nil, err
}
