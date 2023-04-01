package errors

import "fmt"

type EntityError string

func (ee EntityError) Error() string {
	return string(ee)
}

const (
	ErrNullID           = EntityError("entity has null ID")
	ErrInvalidID        = EntityError("entity has invalid ID")
	ErrNullReference    = EntityError("entity has null reference ID")
	ErrInvalidReference = EntityError("entity has invalid reference ID")
	ErrInvalidContent   = EntityError("entity contains invalid content")
	ErrInvalidDate      = EntityError("entity contains invalid date")
	ErrInvalidName      = EntityError("entity contains invalid name")
	ErrInvalidPassword  = EntityError("entity contains invalid password")
	ErrInvalidRole      = EntityError("entity contains invalid role")
)

func ErrFeatureNotExistWrap(feature string) EntityError {
	return EntityError(fmt.Sprintf("feature %s does not exist", feature))
}
