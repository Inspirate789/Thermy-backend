package postgres_storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PropertiesPgRepository struct {
	conn *sqlx.DB
}

func (r *PropertiesPgRepository) GetAllProperties() ([]entities.Property, error) {
	var properties []entities.Property
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &properties, selectAllProperties, nil)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *PropertiesPgRepository) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	args := map[string]any{
		"layer_name": layer,
		"lang":       unit.Lang,
		"unit_text":  unit.Text,
	}
	var properties []entities.Property
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &properties, selectPropertiesByUnit, args)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *PropertiesPgRepository) SaveProperties(properties []string) ([]int, error) {
	args := map[string]any{
		"properties_array": pq.Array(properties),
	}
	var id []int
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, &id, insertProperties, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
