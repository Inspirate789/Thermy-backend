package errors

type AuthorizationError string

func (ae AuthorizationError) Error() string {
	return string(ae)
}

const (
	// ErrInvalidToken  = AuthorizationError("invalid token")
	ErrOpenSession           = AuthorizationError("cannot add new session")
	ErrRemoveDatabaseSession = AuthorizationError("cannot remove session: internal database error")
	ErrRemoveSessionByToken  = AuthorizationError("cannot remove session: invalid token")
	ErrGetSession            = AuthorizationError("cannot find session: invalid token")
)
