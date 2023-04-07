package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type ModelsPgRepository struct{}

func (r *ModelsPgRepository) GetAllModels(conn storage.ConnDB, layer string) ([]entities.Model, error) {
	args := map[string]interface{}{
		"layer_name": layer,
	}
	return namedSelectSliceFromScript[[]entities.Model](conn, selectAllModelsQuery, args)
}

func (r *ModelsPgRepository) SaveModels(conn storage.ConnDB, layer string, models []string) ([]int, error) {
	args := map[string]interface{}{
		"layer_name":   layer,
		"models_array": pq.Array(models),
	}
	return namedSelectSliceFromScript[[]int](conn, insertModelsQuery, args)
}
