package delivery

import (
	"book-service/src/common/log"
	"book-service/src/infrastructure/cbreaker"
	"book-service/src/infrastructure/config"
	"book-service/src/interface/delivery"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book_category"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookCategoryGrpcImpl struct {
	client pb.BookCategoryServiceClient
}

func NewBookCategoryGrpc() (delivery.BookCategoryGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.NewClient(config.Conf.BookCategoryService.GrpcAddress, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.NewBookCategoryGrpc", "section": "grpc.NewClient"}).Fatal(err)
	}

	client := pb.NewBookCategoryServiceClient(conn)

	return &BookCategoryGrpcImpl{
		client: client,
	}, conn
}

func (d *BookCategoryGrpcImpl) DeleteBookFromCategoryReq(ctx context.Context, bookId int) error {
	_, err := cbreaker.BookCategoryGrpc.Execute(func() (any, error) {
		_, err := d.client.DeleteBookFromCategoryReq(ctx, &pb.BookId{Id: uint32(bookId)})
		return nil, err
	})

	return err
}