package middleware

import "fmt"

type MiddlewareError string

func (e MiddlewareError) Error() string {
	return string(e)
}

func ErrUserNotExist(token string) MiddlewareError {
	return MiddlewareError(fmt.Sprintf("user with token %s does not exist", token))
}

func ErrInvalidRole(expected, got string) MiddlewareError {
	return MiddlewareError(fmt.Sprintf("invalid role: expected %s, got %s", expected, got))
}
