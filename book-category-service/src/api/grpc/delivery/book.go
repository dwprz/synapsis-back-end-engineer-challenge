package delivery

import (
	"book-category-service/src/common/log"
	"book-category-service/src/infrastructure/cbreaker"
	"book-category-service/src/infrastructure/config"
	"book-category-service/src/interface/delivery"
	"context"
	"fmt"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookGrpcImpl struct {
	client pb.BookServiceClient
}

func NewBookGrpc() (delivery.BookGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.NewClient(config.Conf.BookService.GrpcAddress, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.NewBookGrpc", "section": "grpc.NewClient"}).Fatal(err)
	}

	client := pb.NewBookServiceClient(conn)

	return &BookGrpcImpl{
		client: client,
	}, conn
}

func (d *BookGrpcImpl) FindManyByIds(ctx context.Context, bookIds []int) ([]*pb.Book, error) {
	var ids []uint32
	if err := copier.Copy(&ids, bookIds); err != nil {
		return nil, err
	}

	res, err := cbreaker.BookGrpc.Execute(func() (any, error) {
		res, err := d.client.FindManyByIds(ctx, &pb.BookIds{
			Ids: ids,
		})

		return res.Data, err
	})

	if err != nil {
		return nil, err
	}

	book, ok := res.([]*pb.Book)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T ([]*pb.Book)", book)
	}

	return book, nil
}