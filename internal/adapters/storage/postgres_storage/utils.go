package postgres_storage

import (
	"context"
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
)

func executeScript(conn storage.ConnDB, script string, args ...any) error {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return errors.New("cannot get *sqlx.DB from argument")
	}

	_, err := sqlx_utils.Exec(context.Background(), sqlxDB, script, args...)

	return err
}

func executeNamedScript(conn storage.ConnDB, script string, args map[string]any) error {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return errors.New("cannot get *sqlx.DB from argument")
	}

	_, err := sqlx_utils.NamedExec(context.Background(), sqlxDB, script, args)

	return err
}

func selectValueFromScript[T any](conn storage.ConnDB, script string, args ...any) (T, error) {
	var value T

	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return value, errors.New("cannot get *sqlx.DB from argument")
	}

	err := sqlx_utils.Get(context.Background(), sqlxDB, &value, script, args...)
	if err != nil {
		return value, err
	}

	return value, nil
}

func selectValueFromNamedScript[T any](conn storage.ConnDB, script string, args map[string]any) (T, error) {
	var value T

	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return value, errors.New("cannot get *sqlx.DB from argument")
	}

	err := sqlx_utils.NamedGet(context.Background(), sqlxDB, &value, script, args)
	if err != nil {
		return value, err
	}

	return value, nil
}

func selectSliceFromScript[S ~[]E, E any](conn storage.ConnDB, script string, args ...any) (S, error) {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return nil, errors.New("cannot get *sqlx.DB from argument") // TODO: log and wrap
	}

	var slice S
	err := sqlx_utils.Select(context.Background(), sqlxDB, &slice, script, args...)
	if err != nil {
		return nil, err // TODO: log and wrap
	}

	return slice, nil
}

func namedSelectSliceFromScript[S ~[]E, E any](conn storage.ConnDB, script string, args map[string]any) (S, error) {
	sqlxDB, ok := conn.(*sqlx.DB)
	if !ok {
		return nil, errors.New("cannot get *sqlx.DB from argument") // TODO: log and wrap
	}

	var slice S
	err := sqlx_utils.NamedSelect(context.Background(), sqlxDB, &slice, script, args)
	if err != nil {
		return nil, err // TODO: log and wrap
	}

	return slice, nil
}
