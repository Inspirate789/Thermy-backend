package interfaces

import (
	"context"
)

type AuthRequest struct {
	Username, Password string
}

type Storage interface {
	Connect(*AuthRequest, context.Context) (string, error) // Get role in database and error
}
