package errors

type AuthorizationError string

func (ae AuthorizationError) Error() string {
	return string(ae)
}

const (
	ErrInvalidToken  = AuthorizationError("invalid token")
	ErrRemoveSession = AuthorizationError("cannot remove session: invalid token")
	ErrGetSession    = AuthorizationError("cannot find session: invalid token")
)
