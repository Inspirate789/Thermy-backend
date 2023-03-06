package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type LayersPgRepository struct{}

func (r *LayersPgRepository) LayerExist(conn storage.ConnDB, layer string) (bool, error) {
	layers, err := selectSliceFromScript[[]string](conn, "sql/select_layer_names.sql")
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
	return selectSliceFromScript[[]string](conn, "sql/select_layer_names.sql")
}

func (r *LayersPgRepository) SaveLayer(conn storage.ConnDB, name string) error {
	return executeScript(conn, "sql/create_layer.sql", name)
}
