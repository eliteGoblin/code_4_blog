package action

import (
	"net/http"
	"fmt"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "hello Fairfax")
}