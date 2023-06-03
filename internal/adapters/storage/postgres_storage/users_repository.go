package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type UsersPgRepository struct{}

func (r *UsersPgRepository) AddUser(conn storage.ConnDB, user interfaces.UserDTO) error {
	args := map[string]any{
		"username": user.Name,
		"password": user.Password,
		"role":     user.Role,
	}

	return executeNamedScript(conn, insertUser, args)
}

//func (r *UsersPgRepository) GetUserPassword(conn storage.ConnDB, username string) (string, error) {
//	return selectValueFromScript[string](conn, selectUserPassword, username)
//}
