package postgres_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage/wrappers"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/jmoiron/sqlx"
	"os"
)

// Execute from project root (backend/): go-bindata -pkg postgres_storage -o internal/adapters/storage/postgres_storage/sqlscripts.go ./internal/adapters/storage/postgres_storage/sql
// or execute in GoLand
//go:generate go-bindata -pkg postgres_storage -o sqlscripts.go ./sql

type PostgresStorage struct{}

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

func getPostgresInfo(request *interfaces.AuthRequest) (string, error) {
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

	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		request.Username,
		request.Password,
		dbName,
		sslMode)

	return postgresInfo, nil
}

func (ps *PostgresStorage) OpenConn(request *interfaces.AuthRequest, ctx context.Context) (interfaces.ConnDB, string, error) {
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

func (ps *PostgresStorage) CloseConn(db interfaces.ConnDB) error {
	sqlxDB, ok := db.(*sqlx.DB)
	if !ok {
		return errors.New("cannot get *sqlx.DB from argument")
	}

	return sqlxDB.Close()
}
