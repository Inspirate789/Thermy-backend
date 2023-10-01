package postgres_storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	conn *sqlx.DB
	LayersPgRepository
	ModelsPgRepository
	ModelElementsPgRepository
	PropertiesPgRepository
	UnitsPgRepository
	UsersPgRepository
}

func NewPostgresStorage(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{
		conn:                      db,
		LayersPgRepository:        LayersPgRepository{conn: db},
		ModelsPgRepository:        ModelsPgRepository{conn: db},
		ModelElementsPgRepository: ModelElementsPgRepository{conn: db},
		PropertiesPgRepository:    PropertiesPgRepository{conn: db},
		UnitsPgRepository:         UnitsPgRepository{conn: db},
		UsersPgRepository:         UsersPgRepository{conn: db},
	}
}
