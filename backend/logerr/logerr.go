package logerr

import (
	"fmt"

	"log"
)

// Ideally would have some form of logging to an APM here, but trying to stay within scope of challenge.
func logRed(str string) {
	log.Printf("\x1b[31;1m🐞 %+v\x1b[0m", str)
}

func FromError(err error) {
	logRed(fmt.Sprintf("Error: %+v", err))
}

func FromPanic(err interface{}) {
	logRed(fmt.Sprintf("Panic: %+v", err))
}

func FromHTTPRequestError(
	reportedErr error,
) {
	FromError(reportedErr)
}

func FromHTTPRequestPanic(
	reportedErr interface{},
) {
	FromPanic(fmt.Sprintf("Panic: %+v", reportedErr))
}
