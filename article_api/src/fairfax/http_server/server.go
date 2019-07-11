package http_server

import (
	"fmt"
	"net/http"
)

var (
	articleApiAddr = "0.0.0.0"
	articleApiPort = "8080"
)
func ListenAndServe() {
	router := NewRouter()

	addr := fmt.Sprintf("%s:%s", articleApiAddr, articleApiPort)
	http.ListenAndServe(addr, router)
}

