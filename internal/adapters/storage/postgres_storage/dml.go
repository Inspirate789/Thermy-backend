package postgres_storage

import (
	"context"
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage/wrappers"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/jmoiron/sqlx"
	"github.com/sethvargo/go-password/password"
)

type JoinedUnits struct {
	UnitRuID      int    `db:"unit_ru_id"`
	UnitRuModelID int    `db:"unit_ru_model_id"`
	UnitRuRegDate string `db:"unit_ru_registration_date"`
	UnitRuText    string `db:"unit_ru_text"`
	UnitEnID      int    `db:"unit_en_id"`
	UnitEnModelID int    `db:"unit_en_model_id"`
	UnitEnRegDate string `db:"unit_en_registration_date"`
	UnitEnText    string `db:"unit_en_text"`
}

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

func (ps *PostgresStorage) GetAllModels(conn storage.ConnDB, layer string) ([]entities.Model, error) {
	return selectSliceFromScript[[]entities.Model](conn, "sql/select_all_models.sql", layer)
}

func (ps *PostgresStorage) SaveModels(conn storage.ConnDB, layer string, models []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveModels") // TODO: implement me
}

func (ps *PostgresStorage) GetAllModelElements(conn storage.ConnDB, layer string) ([]entities.ModelElement, error) {
	return selectSliceFromScript[[]entities.ModelElement](conn, "sql/select_all_model_elements.sql", layer)
}

func (ps *PostgresStorage) SaveModelElements(conn storage.ConnDB, layer string, modelElements []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveModelElements") // TODO: implement me
}

func (ps *PostgresStorage) GetAllProperties(conn storage.ConnDB) ([]entities.Property, error) {
	return selectSliceFromScript[[]entities.Property](conn, "sql/select_all_properties.sql")
}

func (ps *PostgresStorage) GetPropertiesByUnit(conn storage.ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	return selectSliceFromScript[[]entities.Property](conn, "sql/select_properties_by_unit.sql", layer, unit.Lang, unit.Text)
}

func (ps *PostgresStorage) SaveProperties(conn storage.ConnDB, properties []string) ([]int, error) {
	return nil, errors.New("postgres storage does not support function SaveProperties") // TODO: implement me
}

func makeOutputUnitDTO(conn storage.ConnDB, layer string, lang string, unit entities.Unit) (interfaces.OutputUnitDTO, error) {
	propertiesID, err := selectSliceFromScript[[]int](conn, "sql/select_properties_id_by_unit.sql", layer, lang, unit.ID)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	contextsID, err := selectSliceFromScript[[]int](conn, "sql/select_contexts_id_by_unit.sql", layer, lang, unit.ID)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	unitDTO := interfaces.OutputUnitDTO{
		ModelID:      unit.ModelID,
		RegDate:      unit.RegDate,
		Text:         unit.Text,
		PropertiesID: propertiesID,
		ContextsID:   contextsID,
	}

	return unitDTO, nil
}

func (ps *PostgresStorage) GetAllUnits(conn storage.ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	unlinkedUnitsRu, err := selectSliceFromScript[[]entities.Unit](conn, "sql/select_unlinked_units_by_lang.sql", layer, "ru")
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	unlinkedUnitsEn, err := selectSliceFromScript[[]entities.Unit](conn, "sql/select_unlinked_units_by_lang.sql", layer, "en")
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	linkedUnits, err := selectSliceFromScript[[]JoinedUnits](conn, "sql/select_linked_units.sql", layer)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	combinedUnits := make([]map[string]interfaces.OutputUnitDTO, len(unlinkedUnitsRu)+len(unlinkedUnitsEn)+len(linkedUnits))
	uniqueContextsID := make(map[int]bool)

	i := 0
	for _, unit := range unlinkedUnitsRu {
		unitDTO, err := makeOutputUnitDTO(conn, layer, "ru", unit)
		if err != nil {
			return interfaces.OutputUnitsDTO{}, err
		}

		for _, contextID := range unitDTO.ContextsID {
			uniqueContextsID[contextID] = true
		}

		combinedUnits[i]["ru"] = unitDTO
		i++
	}
	for _, unit := range unlinkedUnitsEn {
		unitDTO, err := makeOutputUnitDTO(conn, layer, "en", unit)
		if err != nil {
			return interfaces.OutputUnitsDTO{}, err
		}

		for _, contextID := range unitDTO.ContextsID {
			uniqueContextsID[contextID] = true
		}

		combinedUnits[i]["en"] = unitDTO
		i++
	}
	for _, unitPair := range linkedUnits {
		unitRu := entities.Unit{
			ID:      unitPair.UnitRuID,
			ModelID: unitPair.UnitRuModelID,
			RegDate: unitPair.UnitRuRegDate,
			Text:    unitPair.UnitRuText,
		}
		unitEn := entities.Unit{
			ID:      unitPair.UnitEnID,
			ModelID: unitPair.UnitEnModelID,
			RegDate: unitPair.UnitEnRegDate,
			Text:    unitPair.UnitEnText,
		}

		unitRuDTO, err := makeOutputUnitDTO(conn, layer, "ru", unitRu)
		if err != nil {
			return interfaces.OutputUnitsDTO{}, err
		}
		unitEnDTO, err := makeOutputUnitDTO(conn, layer, "en", unitEn)
		if err != nil {
			return interfaces.OutputUnitsDTO{}, err
		}

		for _, contextID := range unitRuDTO.ContextsID {
			uniqueContextsID[contextID] = true
		}
		for _, contextID := range unitEnDTO.ContextsID {
			uniqueContextsID[contextID] = true
		}

		combinedUnits[i]["ru"] = unitRuDTO
		combinedUnits[i]["en"] = unitEnDTO
		i++
	}

	contextsID := make([]int, 0, len(uniqueContextsID))
	for id := range uniqueContextsID {
		contextsID = append(contextsID, id)
	}

	contexts, err := selectSliceFromScript[[]interfaces.ContextDTO](conn, "sql/select_contexts_by_id.sql", contextsID)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{Units: combinedUnits, Contexts: contexts}, nil
}

func (ps *PostgresStorage) GetUnitsByModels(conn storage.ConnDB, layer string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, errors.New("postgres storage does not support function GetUnitsByModels") // TODO: implement me
}

func (ps *PostgresStorage) GetUnitsByProperties(conn storage.ConnDB, layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, errors.New("postgres storage does not support function GetUnitsByProperties") // TODO: implement me
}

func (ps *PostgresStorage) SaveUnits(conn storage.ConnDB, layer string, data interfaces.SaveUnitsDTO) error {
	return errors.New("postgres storage does not support function SaveUnits") // TODO: implement me
}

func (ps *PostgresStorage) RenameUnit(conn storage.ConnDB, layer string, oldName string, newName string) error {
	return errors.New("postgres storage does not support function RenameUnit") // TODO: implement me
}

func (ps *PostgresStorage) SetUnitProperties(conn storage.ConnDB, layer string, unitName string, propertiesID []int) error {
	return errors.New("postgres storage does not support function SetUnitProperties") // TODO: implement me
}

func (ps *PostgresStorage) LayerExist(conn storage.ConnDB, layer string) (bool, error) {
	layers, err := selectSliceFromScript[[]string](conn, "sql/select_layer_names.sql")
	if err != nil {
		return false, err
	}

	for _, elem := range layers {
		if elem == layer {
			return true, nil
		}
	}
	return false, nil
}

func (ps *PostgresStorage) GetAllLayers(conn storage.ConnDB) ([]string, error) {
	return selectSliceFromScript[[]string](conn, "sql/select_layer_names.sql")
}

func (ps *PostgresStorage) SaveLayer(conn storage.ConnDB, name string) error {
	return executeScript(conn, "sql/create_layer.sql", name)
}

func (ps *PostgresStorage) AddUser(conn storage.ConnDB, username string, role string) error {
	passwd, err := password.Generate(20, 10, 0, false, false)
	if err != nil {
		return err
	}

	return executeScript(conn, "sql/insert_user.sql", username, passwd, role)
}

func (ps *PostgresStorage) GetUserPassword(conn storage.ConnDB, username string) (string, error) {
	return selectValueFromScript[string](conn, "sql/select_user_password.sql", username)
}
