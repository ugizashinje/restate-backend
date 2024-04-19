package messages

import (
	"fmt"
	"net/http"
)

type Error struct {
	status  int    `json:"status"`
	message string `json:"message"`
}

func (e Error) Status() int {
	return e.status
}

func (e Error) Error() string {
	return e.message
}

func Errorf(status int, format string, args ...any) Error {
	e := Error{}
	e.status = status
	e.message = fmt.Sprintf(format, args...)
	return e
}

func Unauthorized() Error {
	return Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation")
}
