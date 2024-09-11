package client

import (
	"book-service/src/common/log"
	"book-service/src/interface/delivery"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// this main grpc client
type Grpc struct {
	BookCategory     delivery.BookCategoryGrpc
	bookCategoryConn *grpc.ClientConn
}

func NewGrpc(bg delivery.BookCategoryGrpc, bookCategoryConn *grpc.ClientConn) *Grpc {
	return &Grpc{
		BookCategory:     bg,
		bookCategoryConn: bookCategoryConn,
	}
}

func (g *Grpc) Close() {
	if err := g.bookCategoryConn.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "client.Grpc/Close", "section": "bookCategoryConn.Close"}).Error(err)
	}
}
