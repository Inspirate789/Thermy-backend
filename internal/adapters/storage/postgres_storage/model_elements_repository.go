package postgres_storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ModelElementsPgRepository struct {
	conn *sqlx.DB
}

func (r *ModelElementsPgRepository) GetAllModelElements(layer string) ([]entities.ModelElement, error) {
	args := map[string]any{
		"layer_name": layer,
	}
	var elements []entities.ModelElement
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, elements, selectAllModelElements, args)
	if err != nil {
		return nil, err
	}
	return elements, nil
}

func (r *ModelElementsPgRepository) SaveModelElements(layer string, modelElements []string) ([]int, error) {
	args := map[string]any{
		"layer_name":     layer,
		"elements_array": pq.Array(modelElements),
	}
	var id []int
	err := sqlx_utils.NamedSelect(context.Background(), r.conn, id, insertModelElements, args)
	if err != nil {
		return nil, err
	}
	return id, nil
}
