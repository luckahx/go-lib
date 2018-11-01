//Package cerr is a wrappable error with json and grpc marshalling
package cerr

//Internal errors
func NewInternal(code string, message string) CError {
	return Wrap(New(code, message), "internal", "internal error")
}
func NewInternalf(code string, messageformat string, vals ...interface{}) CError {
	return Wrap(Newf(code, messageformat, vals...), "internal", "internal error")
}
func WrapInternal(err error, code string, message string) CError {
	return Wrap(Wrap(err, code, message), "internal", "internal error")
}
func WrapInternalf(err error, code string, messageformat string, vals ...interface{}) CError {
	return Wrap(Wrapf(err, code, messageformat, vals...), "internal", "internal error")
}

// invalid requests
func NewInvalidRequest(code string, message string) CError {
	return Wrap(New(code, message), "invalid_request", "invalid_request")
}
func NewInvalidRequestf(code string, messageformat string, vals ...interface{}) CError {
	return Wrap(Newf(code, messageformat, vals...), "invalid_request", "invalid_request")
}
func WrapInvalidRequest(err error, code string, message string) CError {
	return Wrap(Wrap(err, code, message), "invalid_request", "invalid_request")
}
func WrapInvalidRequestf(err error, code string, messageformat string, vals ...interface{}) CError {
	return Wrap(Wrapf(err, code, messageformat, vals...), "invalid_request", "invalid_request")
}
