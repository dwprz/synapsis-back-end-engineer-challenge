package errors

import "google.golang.org/grpc/codes"

type Response struct {
	HttpCode int
	GrpcCode codes.Code
	Message  string
}

func (r *Response) Error() string {
	return r.Message
}
