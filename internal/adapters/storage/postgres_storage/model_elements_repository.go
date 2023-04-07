package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type ModelElementsPgRepository struct{}

func (r *ModelElementsPgRepository) GetAllModelElements(conn storage.ConnDB, layer string) ([]entities.ModelElement, error) {
	args := map[string]interface{}{
		"layer_name": layer,
	}
	return namedSelectSliceFromScript[[]entities.ModelElement](conn, selectAllModelElementsQuery, args)
}

func (r *ModelElementsPgRepository) SaveModelElements(conn storage.ConnDB, layer string, modelElements []string) ([]int, error) {
	args := map[string]interface{}{
		"layer_name":     layer,
		"elements_array": pq.Array(modelElements),
	}
	return namedSelectSliceFromScript[[]int](conn, insertModelElementsQuery, args)
}
