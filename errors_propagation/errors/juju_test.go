package errors

import (
	"fmt"
	"testing"

	"github.com/juju/errors"
)

func TestJuju(t *testing.T) {
	err := errors.New("root cause")
	err = errors.Trace(err)
	err = errors.Annotatef(err, "wrap with annotation\n")
	fmt.Printf("%s", errors.ErrorStack(err))

	fmt.Printf("Cause is %+v\n", errors.Cause(err))
}
