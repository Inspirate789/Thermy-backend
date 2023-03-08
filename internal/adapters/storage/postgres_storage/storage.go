package postgres_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage/wrappers"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

// Execute from current directory: go-bindata -pkg postgres_storage -o sqlscripts.go ./sql
// or execute in GoLand
//go:generate go-bindata -pkg postgres_storage -o sqlscripts.go ./sql

type PostgresStorage struct {
	LayersPgRepository
	ModelsPgRepository
	ModelElementsPgRepository
	PropertiesPgRepository
	UnitsPgRepository
	UsersPgRepository
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

//	AuthDB Example: { // TODO: delete this
//		host:    	"postgres"
//		port:     	5432
//		username: 	"user"
//		password: 	"mypassword"
//		dbname:   	"user_db"
//		sslmode:  	"disable"
//	}

func getConnRole(conn *sqlx.DB, ctx context.Context) (string, error) {
	script, err := Asset("sql/select_role.sql")
	if err != nil {
		return "", err
	}

	var roles []string
	err = wrappers.Select(ctx, conn, &roles, string(script))
	if err != nil {
		return "", err
	}
	if len(roles) != 1 {
		return "", errors.New("")
	}

	return roles[0], nil
}

func getPostgresInfo(request *storage.AuthRequest) (string, error) {
	host := os.Getenv("POSTGRES_HOST") // TODO: get once?
	if host == "" {
		return "", errors.New("POSTGRES_HOST must be set")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return "", errors.New("POSTGRES_PORT must be set")
	}

	dbName := os.Getenv("POSTGRES_DBNAME")
	if dbName == "" {
		return "", errors.New("POSTGRES_DBNAME must be set")
	}

	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	if sslMode == "" {
		return "", errors.New("POSTGRES_SSL_MODE must be set")
	}

	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		request.Username,
		request.Password,
		dbName,
		sslMode)

	return postgresInfo, nil
}

func (ps *PostgresStorage) OpenConn(request *storage.AuthRequest, ctx context.Context) (storage.ConnDB, string, error) {
	postgresInfo, err := getPostgresInfo(request)
	if err != nil {
		return nil, "", err
	}

	sqlDB, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, "", err
	}

	driverName, exists := os.LookupEnv("POSTGRES_DRIVER_NAME")
	if !exists {
		return nil, "", errors.New("POSTGRES_DRIVER_NAME must be set")
	}
	sqlxDB := sqlx.NewDb(sqlDB, driverName) // "postgres" // TODO: delete comment

	role, err := getConnRole(sqlxDB, ctx)
	if err != nil {
		return nil, "", err
	}

	return sqlxDB, role, nil
}

func (ps *PostgresStorage) CloseConn(db storage.ConnDB) error {
	sqlxDB, ok := db.(*sqlx.DB)
	if !ok {
		return errors.New("cannot get *sqlx.DB from argument")
	}

	return sqlxDB.Close()
}
