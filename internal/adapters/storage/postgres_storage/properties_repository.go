package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/models"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type PropertiesPgRepository struct{}

func (r *PropertiesPgRepository) GetAllProperties(conn storage.ConnDB) ([]models.Property, error) {
	return selectSliceFromScript[[]models.Property](conn, "sql/select_all_properties.sql")
}

func (r *PropertiesPgRepository) GetPropertiesByUnit(conn storage.ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]models.Property, error) {
	return selectSliceFromScript[[]models.Property](conn, "sql/select_properties_by_unit.sql", layer, unit.Lang, unit.Text)
}

func (r *PropertiesPgRepository) SaveProperties(conn storage.ConnDB, properties []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveProperties") // TODO: implement me
}
