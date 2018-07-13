package errors

import (
	"fmt"
)

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func MissingRequiredCommandFlag(command string, flag string) *Error {
	return &Error{
		message: fmt.Sprintf("Missing required flag %s for command %s", flag, command),
	}
}

func DuplicateHandlerConfig(handler string) *Error {
	return &Error{
		message: fmt.Sprintf("Duplicate handler config specified in config file : %s", handler),
	}
}
