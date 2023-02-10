package httperr

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int
}

func HTTPErr(c int) HTTPError {
	return HTTPError{StatusCode: c}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("Http Error %d: %s", e.StatusCode, http.StatusText(e.StatusCode))
}

type HumanReadableErr struct {
	msg string
}

func (e HumanReadableErr) Error() string {
	return e.msg
}

func HumanReadable(err string) error {
	return HumanReadableErr{
		msg: err,
	}
}
