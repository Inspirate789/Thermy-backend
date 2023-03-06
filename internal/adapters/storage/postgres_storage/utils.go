package postgres_storage

import (
	"context"
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage/wrappers"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/jmoiron/sqlx"
)

func executeScript(conn storage.ConnDB, scriptName string, args ...any) error {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return errors.New("cannot get *sqlx.DB from argument")
	}

	script, err := Asset(scriptName)
	if err != nil {
		return err
	}

	_, err = wrappers.Exec(context.Background(), sqlxDB, string(script), args)

	return err
}

func selectValueFromScript[T any](conn storage.ConnDB, scriptName string, args ...any) (T, error) {
	var value T

	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return value, errors.New("cannot get *sqlx.DB from argument")
	}

	script, err := Asset(scriptName)
	if err != nil {
		return value, err
	}

	err = wrappers.Get(context.Background(), sqlxDB, &value, string(script), args)
	if err != nil {
		return value, err
	}

	return value, nil
}

func selectSliceFromScript[S ~[]E, E any](conn storage.ConnDB, scriptName string, args ...any) (S, error) {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return nil, errors.New("cannot get *sqlx.DB from argument") // TODO: log and wrap
	}

	script, err := Asset(scriptName)
	if err != nil {
		return nil, err // TODO: log and wrap
	}

	var slice S
	err = wrappers.Select(context.Background(), sqlxDB, &slice, string(script), args)
	if err != nil {
		return nil, err // TODO: log and wrap
	}

	return slice, nil
}
