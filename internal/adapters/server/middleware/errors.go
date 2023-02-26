package middleware

import "fmt"

type MiddlewareError string

func (e MiddlewareError) Error() string {
	return string(e)
}

func ErrInvalidRole(expected, got string) MiddlewareError {
	return MiddlewareError(fmt.Sprintf("invalid role: expected %s, got %s", expected, got))
}

//const (
//	ErrInvalidRole = MiddlewareError("invalid token")
//)
