package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type ModelsPgRepository struct{}

func (r *ModelsPgRepository) GetAllModels(conn storage.ConnDB, layer string) ([]entities.Model, error) {
	args := map[string]interface{}{
		"layer_name": layer,
	}
	return namedSelectSliceFromScript[[]entities.Model](conn, selectAllModelsQuery, args)
}

func (r *ModelsPgRepository) SaveModels(conn storage.ConnDB, layer string, models []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveModels") // TODO: implement me
}
