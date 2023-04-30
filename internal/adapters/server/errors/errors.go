package errors

import "fmt"

type ServerError string

func (se ServerError) Error() string {
	return string(se)
}

func ErrCannotParseJSONWrap(entityName string) ServerError {
	return ServerError(fmt.Sprintf("cannot parse %s from received JSON", entityName))
}

func ErrCannotParseURLWrap(param string) ServerError {
	return ServerError(fmt.Sprintf("cannot parse %s from requestURL", param))
}

const (
	ErrAccessSystemInfo = ServerError("failed to access system information")
)
