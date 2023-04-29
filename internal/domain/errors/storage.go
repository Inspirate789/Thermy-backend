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
	ErrAuthDatabase    = StorageError("authentication failed")
	ErrEnvironment     = StorageError("incomplete environment")
)

func ErrUnknownLayerWrap(layer string) StorageError {
	return StorageError(fmt.Sprintf("layer %s does not exist in database", layer))
}

func IdentifyStorageError(err error) StorageError {
	switch {
	case strings.Contains(strings.ToLower(err.Error()), "authentication failed"):
		return ErrConnectDatabase
	case strings.Contains(strings.ToLower(err.Error()), "bad connection"):
		return ErrAuthDatabase
	case strings.Contains(strings.ToLower(err.Error()), "must be set"):
		return ErrEnvironment
	default:
		return ErrDatabase
	}
}
