package werror

import (
	"fmt"
	"log"
	"sync"
)

type Error struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"detail"`
}

func (e Error) Error() string {
	if e.Code == "" {
		return e.Message
	}
	if len(e.Details) > 0 {
		return fmt.Sprintf("%s - %v", e.Code, e.Details)
	}
	return fmt.Sprintf("%s - %v", e.Code, e.Message)
}

func (e Error) AttacthDetail(detail map[string]any) Error {
	e.Details = detail
	return e
}

func (e Error) Is(err error) bool {
	myErr, ok := err.(Error)
	return ok && e.Code == myErr.Code
}

type Errors struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`

	mut *sync.Mutex
}

func (e Errors) Error() string {
	errLength := len(e.Errors)
	if errLength == 0 {
		return ""
	}
	message := fmt.Sprintf("%v: %v,\n", "message", e.Message)
	result := message + "errors:["
	for i, v := range e.Errors {
		separator := ""
		if i < errLength-1 {
			separator = ", "
		}
		result += v.Error() + separator
	}
	result += "]"
	return result
}

func NewError(message string) *Errors {
	return &Errors{
		Message: message,
		Errors:  make([]Error, 0),
		mut:     &sync.Mutex{},
	}
}

func (e *Errors) Add(err error) {
	e.mut.Lock()

	switch errors := err.(type) {
	case Error:
		e.Errors = append(e.Errors, errors)
	case *Error:
		e.Errors = append(e.Errors, *errors)
	case Errors:
		e.Errors = append(e.Errors, errors.Errors...)
	case *Errors:
		e.Errors = append(e.Errors, errors.Errors...)
	case error:
		e.Errors = append(e.Errors, Error{Message: errors.Error()})
	default:
		log.Fatal("invalid error, not supported type")
	}

	e.mut.Unlock()
}

func (e *Errors) Return() error {
	if len(e.Errors) > 0 {
		return e
	}
	return nil
}
