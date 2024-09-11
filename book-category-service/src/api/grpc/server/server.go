package server

import (
	"book-category-service/src/api/grpc/interceptor"
	"book-category-service/src/common/log"
	"net"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book_category"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Grpc struct {
	Address                  string
	server                   *grpc.Server
	bookCategoryGrpcHandler  pb.BookCategoryServiceServer
	unaryResponseInterceptor *interceptor.UnaryResponse
}

// this main grpc server
func NewGrpc(address string, bh pb.BookCategoryServiceServer, ur *interceptor.UnaryResponse) *Grpc {
	return &Grpc{
		Address:                  address,
		bookCategoryGrpcHandler:  bh,
		unaryResponseInterceptor: ur,
	}
}

func (g *Grpc) Run() {
	listener, err := net.Listen("tcp", g.Address)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "server.Grpc/Run", "section": "net.Listen"}).Fatal(err)
	}

	log.Logger.Infof("grpc run in: %s", g.Address)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			g.unaryResponseInterceptor.Recovery,
			g.unaryResponseInterceptor.Error,
		))

	g.server = grpcServer

	pb.RegisterBookCategoryServiceServer(grpcServer, g.bookCategoryGrpcHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "server.Grpc/Run", "section": "grpcServer.Serve"}).Fatal(err)
	}
}

func (g *Grpc) Stop() {
	g.server.Stop()
}
