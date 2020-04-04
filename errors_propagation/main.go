package main

import (
	"error_propagation/mware"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/users", mware.AddErrorHandler(handleUser))

	http.ListenAndServe(":8090", nil)
}
