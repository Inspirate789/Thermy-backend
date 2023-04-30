package errors

import "fmt"

type AuthorizationError string

func (ae AuthorizationError) Error() string {
	return string(ae)
}

const (
	// ErrInvalidToken  = AuthorizationError("invalid token")
	ErrRemoveSessionByToken = AuthorizationError("cannot remove session: invalid token")
	ErrGetSession           = AuthorizationError("cannot find session: invalid token")
)

func ErrOpenSessionWrap(err error) AuthorizationError {
	return AuthorizationError(fmt.Sprintf("cannot open new session: %v", err))
}

func ErrCloseDatabaseSessionWrap(err error) AuthorizationError {
	return AuthorizationError(fmt.Sprintf("cannot close session: internal database error: %v", err))
}
