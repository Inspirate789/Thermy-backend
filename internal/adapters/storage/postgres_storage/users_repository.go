package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/sethvargo/go-password/password"
)

type UsersPgRepository struct{}

func (r *UsersPgRepository) AddUser(conn storage.ConnDB, username string, role string) error {
	_, err := password.Generate(20, 10, 0, false, false) // TODO: take from frontend
	if err != nil {
		return err
	}

	args := map[string]interface{}{
		"username": username,
		"password": "passwd",
		"role":     role,
	}

	return executeNamedScript(conn, insertUserQuery, args)
}

func (r *UsersPgRepository) GetUserPassword(conn storage.ConnDB, username string) (string, error) { // TODO: remove
	return selectValueFromScript[string](conn, selectUserPasswordQuery, username)
}
