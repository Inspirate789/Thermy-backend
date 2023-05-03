package errors

import (
	"fmt"
	"strings"
)

type StorageError string

func (se StorageError) Error() string {
	return string(se)
}

const (
	ErrDatabase        = StorageError("internal database error")
	ErrConnectDatabase = StorageError("bad connection to the database")
	ErrDuplicateData   = StorageError("the data to be added already exists in the database")
	ErrPermission      = StorageError("not enough permissions to access the data")
	ErrAuthDatabase    = StorageError("authentication failed")
	ErrEnvironment     = StorageError("incomplete environment")
)

func ErrUnknownLayerWrap(layer string) StorageError {
	return StorageError(fmt.Sprintf("layer %s does not exist in database", layer))
}

func IdentifyStorageError(err error) StorageError {
	switch {
	case strings.Contains(strings.ToLower(err.Error()), "authentication failed"):
		return ErrAuthDatabase
	case strings.Contains(strings.ToLower(err.Error()), "bad connection"):
		return ErrConnectDatabase
	case strings.Contains(strings.ToLower(err.Error()), "permission denied"):
		return ErrPermission
	case strings.Contains(strings.ToLower(err.Error()), "already exist"):
		return ErrDuplicateData
	case strings.Contains(strings.ToLower(err.Error()), "must be set"):
		return ErrEnvironment
	default:
		return ErrDatabase
	}
}
