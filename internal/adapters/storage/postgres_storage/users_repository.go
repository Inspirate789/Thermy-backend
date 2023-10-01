package postgres_storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
)

type UsersPgRepository struct {
	conn *sqlx.DB
}

func (r *UsersPgRepository) AddUser(user interfaces.UserDTO) error {
	args := map[string]any{
		"username": user.Name,
		"password": user.Password,
		"role":     user.Role,
	}
	_, err := sqlx_utils.NamedExec(context.Background(), r.conn, insertUser, args)
	return err
}

func (r *UsersPgRepository) GetUser(request entities.AuthRequest) (*entities.User, error) {
	args := map[string]any{
		"username": request.Username,
		"password": request.Password,
	}
	var user entities.User
	err := sqlx_utils.NamedGet(context.Background(), r.conn, &user, selectUser, args)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
