//Package cerr is a wrappable error with json and grpc marshalling
package cerr

import (
	"fmt"
	"time"
)

type CError interface {
	GetCode() string
	GetMsg() string
	GetCause() error
	GetTS() int64

	Error() string

	FullMessage() string
	FullErrorStack() []CError
	Has(code string) CError
	HasFirst(codes ...string) (string, CError)
}

type CErrorMessage struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg"`
	TS   int64  `json:"ts"`

	Cause *CErrorMessage `json:"cause,omitempty"`
}

func (e *CErrorMessage) GetCode() string { return e.Code }
func (e *CErrorMessage) GetMsg() string  { return e.Msg }
func (e *CErrorMessage) GetTS() int64    { return e.TS }
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
	if e == nil {
		return nil
	}
	errs := []CError{e}
	if e.Cause != nil {
		errs = append(errs, e.Cause.FullErrorStack()...)
	}
	return errs
}
func (e *CErrorMessage) Has(code string) CError {
	for _, err := range e.FullErrorStack() {
		if err.GetCode() == code {
			return err
		}
	}
	return nil
}
func (e *CErrorMessage) HasFirst(codes ...string) (string, CError) {
	for _, err := range e.FullErrorStack() {
		for _, code := range codes {
			if err.GetCode() == code {
				return code, err
			}
		}
	}
	return "", nil
}

func New(code string, message string) CError {
	return &CErrorMessage{TS: ts(), Code: code, Msg: message}
}
func Newf(code string, messageformat string, vals ...interface{}) CError {
	return New(code, fmt.Sprintf(messageformat, vals...))
}

func Wrap(err error, code string, message string) CError {
	if err == nil {
		return nil
	}
	return &CErrorMessage{TS: ts(), Code: code, Msg: message, Cause: FromError(err)}
}
func Wrapf(err error, code string, messageformat string, vals ...interface{}) CError {
	if err == nil {
		return nil
	}
	return Wrap(err, code, fmt.Sprintf(messageformat, vals...))
}

func FromError(e error) *CErrorMessage {
	if e == nil {
		return nil
	}
	if fullError, ok := e.(CError); ok {
		res := &CErrorMessage{TS: fullError.GetTS(), Code: fullError.GetCode(), Msg: fullError.GetMsg()}
		if fullError.GetCause() != nil {
			res.Cause = FromError(fullError.GetCause())
		}
		return res
	}
	return &CErrorMessage{TS: ts(), Msg: e.Error()}
}

func ts() int64 {
	return time.Now().UTC().UnixNano()
}
