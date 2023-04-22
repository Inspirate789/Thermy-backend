package postgres_storage

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type PropertiesPgRepository struct{}

func (r *PropertiesPgRepository) GetAllProperties(conn storage.ConnDB) ([]entities.Property, error) {
	return namedSelectSliceFromScript[[]entities.Property](conn, selectAllProperties, make(map[string]any))
}

func (r *PropertiesPgRepository) GetPropertiesByUnit(conn storage.ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	args := map[string]any{
		"layer_name": layer,
		"lang":       unit.Lang,
		"unit_text":  unit.Text,
	}
	fmt.Println(args)
	return namedSelectSliceFromScript[[]entities.Property](conn, selectPropertiesByUnit, args)
}

func (r *PropertiesPgRepository) SaveProperties(conn storage.ConnDB, properties []string) ([]int, error) {
	args := map[string]any{
		"properties_array": pq.Array(properties),
	}
	return namedSelectSliceFromScript[[]int](conn, insertProperties, args)
}
