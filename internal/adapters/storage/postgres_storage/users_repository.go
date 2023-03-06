package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/sethvargo/go-password/password"
)

type UsersPgRepository struct{}

func (r *UsersPgRepository) AddUser(conn storage.ConnDB, username string, role string) error {
	passwd, err := password.Generate(20, 10, 0, false, false)
	if err != nil {
		return err
	}

	return executeScript(conn, "sql/insert_user.sql", username, passwd, role)
}

func (r *UsersPgRepository) GetUserPassword(conn storage.ConnDB, username string) (string, error) {
	return selectValueFromScript[string](conn, "sql/select_user_password.sql", username)
}
