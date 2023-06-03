package errors

import "fmt"

type EntityError string

func (ee EntityError) Error() string {
	return string(ee)
}

func ErrInvalidIdWrap(id int) EntityError {
	return EntityError(fmt.Sprintf("entity has invalid ID: %d", id))
}

func ErrNullIdWrap(id int) EntityError {
	return EntityError(fmt.Sprintf("entity has null ID: %d", id))
}

func ErrInvalidReferenceWrap(id int) EntityError {
	return EntityError(fmt.Sprintf("entity has invalid reference ID: %d", id))
}

func ErrNullReferenceWrap(id int) EntityError {
	return EntityError(fmt.Sprintf("entity has null reference ID: %d", id))
}

func ErrInvalidContentWrap(content any) EntityError {
	return EntityError(fmt.Sprintf("entity contains invalid content: %v", content))
}

func ErrInvalidDateWrap(content any) EntityError {
	return EntityError(fmt.Sprintf("entity contains invalid date: %v", content))
}

func ErrInvalidNameWrap(content any) EntityError {
	return EntityError(fmt.Sprintf("entity contains invalid name: %v", content))
}

func ErrInvalidPasswordWrap(content any) EntityError {
	return EntityError(fmt.Sprintf("entity contains invalid password: %v", content))
}

func ErrInvalidRoleWrap(content any) EntityError {
	return EntityError(fmt.Sprintf("entity contains invalid role: %v", content))
}

func ErrFeatureNotExistWrap(feature string) EntityError {
	return EntityError(fmt.Sprintf("feature %s does not exist", feature))
}
