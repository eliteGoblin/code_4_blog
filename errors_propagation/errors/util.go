package errors

import (
	"context"
	"net/http"

	"github.com/juju/errors"
)

func HTTPStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if errors.Cause(err) == context.Canceled {
		return 499
	}

	if hs, ok := errors.Cause(err).(HasHTTPStatus); ok {
		return hs.HTTPStatus()
	}

	return http.StatusInternalServerError
}

func ResponseHeaders(err error) map[string]string {
	if hasHeader, ok := errors.Cause(err).(HasResponseHeaders); ok {
		return hasHeader.ResponseHeaders()
	}
	return map[string]string{}
}

func IsServiceFailure(err error) bool {
	cause := errors.Cause(err)

	if cause == context.Canceled {
		return false
	}

	if sf, ok := cause.(HasServiceFailure); ok {
		return sf.IsServiceFailure()
	}
	return true
}
