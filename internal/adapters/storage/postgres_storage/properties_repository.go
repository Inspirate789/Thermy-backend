package postgres_storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type PropertiesPgRepository struct{}

func (r *PropertiesPgRepository) GetAllProperties(conn storage.ConnDB) ([]entities.Property, error) {
	return namedSelectSliceFromScript[[]entities.Property](conn, selectAllPropertiesQuery, make(map[string]interface{}))
}

func (r *PropertiesPgRepository) GetPropertiesByUnit(conn storage.ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	args := map[string]interface{}{
		"layer_name": layer,
		"lang":       unit.Lang,
		"unit_text":  unit.Text,
	}
	return namedSelectSliceFromScript[[]entities.Property](conn, selectPropertiesByUnitQuery, args)
}

func (r *PropertiesPgRepository) SaveProperties(conn storage.ConnDB, properties []string) ([]int, error) {
	args := map[string]interface{}{
		"properties_array": pq.Array(properties),
	}
	return namedSelectSliceFromScript[[]int](conn, insertPropertiesQuery, args)
}
