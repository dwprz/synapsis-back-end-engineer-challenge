package client

import (
	"book-category-service/src/common/log"
	"book-category-service/src/interface/delivery"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// this main grpc client
type Grpc struct {
	Book     delivery.BookGrpc
	bookConn *grpc.ClientConn
}

func NewGrpc(bg delivery.BookGrpc, bookConn *grpc.ClientConn) *Grpc {
	return &Grpc{
		Book:     bg,
		bookConn: bookConn,
	}
}

func (g *Grpc) Close() {
	if err := g.bookConn.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "client.Grpc/Close", "section": "bookConn.Close"}).Error(err)
	}
}
