//Package cerr is a wrappable error with json and grpc marshalling
package cerr

import (
	"fmt"
)

type CError interface {
	GetCode() string
	GetMsg() string
	GetCause() error

	Error() string

	FullMessage() string
	FullErrorStack() []CError
}

type CErrorMessage struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg"`

	Cause *CErrorMessage `json:"cause,omitempty"`
}

func (e *CErrorMessage) GetCode() string { return e.Code }
func (e *CErrorMessage) GetMsg() string  { return e.Msg }

func (e *CErrorMessage) GetCause() error {
	if e.Cause != nil {
		return e.Cause
	}
	return nil
}

func (e *CErrorMessage) Error() string {
	if e.Code == "" {
		e.Code = "nocode"
	}
	return fmt.Sprintf("%s:\"%+v\"", e.Code, e.Msg)
}

func (e *CErrorMessage) FullMessage() string {
	s := e.Error()
	if e.Cause != nil {
		s = fmt.Sprintf("%s, %s", s, e.Cause.FullMessage())
	}
	return s
}
func (e *CErrorMessage) FullErrorStack() []CError {
	errs := []CError{e}
	if e.Cause != nil {
		errs = append(errs, e.Cause.FullErrorStack()...)
	}
	return errs
}

func New(code string, message string) *CErrorMessage {
	return &CErrorMessage{Code: code, Msg: message}
}

func Wrap(err error, code string, message string) *CErrorMessage {
	if err == nil {
		return New(code, message)
	}
	return &CErrorMessage{Code: code, Msg: message, Cause: FromError(err)}
}

func FromError(e error) *CErrorMessage {
	if e == nil {
		return nil
	}
	if fullError, ok := e.(CError); ok {
		res := &CErrorMessage{Code: fullError.GetCode(), Msg: fullError.GetMsg()}
		if fullError.GetCause() != nil {
			res.Cause = FromError(fullError.GetCause())
		}
		return res
	}
	return &CErrorMessage{Msg: e.Error()}
}
