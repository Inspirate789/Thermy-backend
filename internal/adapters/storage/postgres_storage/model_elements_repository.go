package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type ModelElementsPgRepository struct{}

func (r *ModelElementsPgRepository) GetAllModelElements(conn storage.ConnDB, layer string) ([]entities.ModelElement, error) {
	return selectSliceFromScript[[]entities.ModelElement](conn, "sql/select_all_model_elements.sql", layer)
}

func (r *ModelElementsPgRepository) SaveModelElements(conn storage.ConnDB, layer string, modelElements []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveModelElements") // TODO: implement me
}
