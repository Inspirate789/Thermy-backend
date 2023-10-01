package postgres_storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ModelsPgRepository struct {
	conn *sqlx.DB
}

func (r *ModelsPgRepository) GetAllModels(layer string) ([]entities.Model, error) {
	args := map[string]any{
		"layer_name": layer,
	}
	var models []entities.Model
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &models, selectAllModels, args)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (r *ModelsPgRepository) SaveModels(layer string, models []string) ([]int, error) {
	args := map[string]any{
		"layer_name":   layer,
		"models_array": pq.Array(models),
	}
	var id []int
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &id, insertModels, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
