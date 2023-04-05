package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type LayersPgRepository struct{}

func (r *LayersPgRepository) LayerExist(conn storage.ConnDB, layer string) (bool, error) {
	layers, err := namedSelectSliceFromScript[[]string](conn, selectLayerNamesQuery, make(map[string]interface{})) // TODO: any?
	if err != nil {
		return false, err
	}

	for _, elem := range layers {
		if elem == layer {
			return true, nil
		}
	}
	return false, nil
}

func (r *LayersPgRepository) GetAllLayers(conn storage.ConnDB) ([]string, error) {
	return namedSelectSliceFromScript[[]string](conn, selectLayerNamesQuery, make(map[string]interface{}))
}

func (r *LayersPgRepository) SaveLayer(conn storage.ConnDB, name string) error {
	args := map[string]interface{}{
		"layer_name": name,
	}
	return executeNamedScript(conn, createLayerQuery, args)
}
