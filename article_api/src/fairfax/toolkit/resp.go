package toolkit

import (
	"net/http"
	"encoding/json"
)

const (
	InvalidParam = "invalid_param"
)

func RenderError(res http.ResponseWriter, errMsg string) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(errMsg))
}

func RenderResponse(res http.ResponseWriter, response interface{}) {
	res.Header().Set("Content-Type", "application/json")
	rsJSON, err := json.Marshal(response)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	} else {
		res.Write(rsJSON)
	}
}