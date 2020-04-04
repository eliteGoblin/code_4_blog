package mware

import (
	"encoding/json"
	"net/http"

	se "error_propagation/errors"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

type WebHandlerFunc func(http.ResponseWriter, *http.Request) error

func AddErrorHandler(handler WebHandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			logrus.Errorf("%+v", errors.ErrorStack(err))
			WriteError(w, err)
		}
	})
}

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	status := se.HTTPStatusCode(err)

	extraHeaders := se.ResponseHeaders(err)
	for k, v := range extraHeaders {
		w.Header().Set(k, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(marshalError(err, status))
}

type ErrResultDTO struct {
	Error string `json:"error,omitempty"`
	Code  string `json:"code,omitempty"`
}

func marshalError(err error, status int) []byte {
	var result ErrResultDTO

	if se.IsServiceFailure(err) {
		result.Error = http.StatusText(status)
	} else {
		result.Error = errors.Cause(err).Error()
	}

	if ce, ok := errors.Cause(err).(se.HasErrorCode); ok {
		result.Code = ce.Code()
	}

	if bytes, err := json.Marshal(result); err == nil {
		return bytes
	}

	return []byte(`
		{
		  "error": "error happened when generating results",
		}
	`)
}
