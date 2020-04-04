package errors

type HasHTTPStatus interface {
	HTTPStatus() int
}

type HasResponseHeaders interface {
	ResponseHeaders() map[string]string
}

type HasServiceFailure interface {
	IsServiceFailure() bool
}

type HasErrorCode interface {
	Code() string
}
