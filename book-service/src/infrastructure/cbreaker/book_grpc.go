package cbreaker

import (
	"book-service/src/common/log"
	"time"

	"github.com/sony/gobreaker/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var BookCategoryGrpc *gobreaker.CircuitBreaker[any]

func init() {
	settings := gobreaker.Settings{
		Name:        "book-category-grpc",
		MaxRequests: 3,
		Interval:    1 * time.Minute,
		Timeout:     15 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {

			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failureRatio >= 0.8
		},
		IsSuccessful: func(err error) bool {
			if err == nil {
				return true
			}

			st, ok := status.FromError(err)
			if !ok {
				return false
			}

			statusCodeSuccess := []codes.Code{
				codes.OK,
				codes.InvalidArgument,
				codes.NotFound,
				codes.Canceled,
			}

			for _, code := range statusCodeSuccess {
				if st.Code() == code {
					return true
				}
			}

			return false
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Logger.Infof("circuit breaker %v from status %v to %v", name, from, to)
		},
	}

	BookCategoryGrpc = gobreaker.NewCircuitBreaker[any](settings)
}
