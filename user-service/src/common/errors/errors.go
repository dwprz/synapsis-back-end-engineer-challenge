package errors

type Response struct {
	HttpCode int
	GrpcCode int
	Message  string
}

func (r *Response) Error() string {
	return r.Message
}
