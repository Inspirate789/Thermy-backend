package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type ModelElementsPgRepository struct{}

func (r *ModelElementsPgRepository) GetAllModelElements(conn storage.ConnDB, layer string) ([]entities.ModelElement, error) {
	args := map[string]interface{}{
		"layer_name": layer,
	}
	return namedSelectSliceFromScript[[]entities.ModelElement](conn, selectAllModelElementsQuery, args)
}

func (r *ModelElementsPgRepository) SaveModelElements(conn storage.ConnDB, layer string, modelElements []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveModelElements") // TODO: implement me
}
