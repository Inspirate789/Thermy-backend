package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type UnitsPgRepository struct{}

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

func makeOutputUnitDTO(conn storage.ConnDB, layer string, lang string, unit entities.Unit) (interfaces.OutputUnitDTO, error) {
	args := map[string]interface{}{
		"layer_name": layer,
		"lang":       lang,
		"unit_id":    unit.ID,
	}
	propertiesID, err := namedSelectSliceFromScript[[]int](conn, selectPropertiesIdByUnitIdQuery, args)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	contextsID, err := namedSelectSliceFromScript[[]int](conn, selectContextsIdByUnitQuery, args)
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

func (r *UnitsPgRepository) GetAllUnits(conn storage.ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	unlinkedUnitsRu, err := namedSelectSliceFromScript[[]entities.Unit](conn, selectUnlinkedUnitsByLangQuery, map[string]interface{}{
		"layer_name": layer,
		"lang":       "ru",
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	unlinkedUnitsEn, err := namedSelectSliceFromScript[[]entities.Unit](conn, selectUnlinkedUnitsByLangQuery, map[string]interface{}{
		"layer_name": layer,
		"lang":       "en",
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	linkedUnits, err := namedSelectSliceFromScript[[]JoinedUnits](conn, selectAllLinkedUnitsQuery, map[string]interface{}{
		"layer_name": layer,
	})
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

	contexts, err := selectSliceFromScript[[]interfaces.ContextDTO](conn, selectContextsByIdQuery, pq.Array(contextsID))
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{Units: combinedUnits, Contexts: contexts}, nil
}

func (r *UnitsPgRepository) GetUnitsByModels(conn storage.ConnDB, layer string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, errors.New("postgres storage does not support function GetUnitsByModels") // TODO: implement me
}

func (r *UnitsPgRepository) GetUnitsByProperties(conn storage.ConnDB, layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, errors.New("postgres storage does not support function GetUnitsByProperties") // TODO: implement me
}

func (r *UnitsPgRepository) SaveUnits(conn storage.ConnDB, layer string, data interfaces.SaveUnitsDTO) error {
	return errors.New("postgres storage does not support function SaveUnits") // TODO: implement me
}

func (r *UnitsPgRepository) RenameUnit(conn storage.ConnDB, layer string, oldName string, newName string) error {
	return errors.New("postgres storage does not support function RenameUnit") // TODO: implement me
}

func (r *UnitsPgRepository) SetUnitProperties(conn storage.ConnDB, layer string, unitName string, propertiesID []int) error {
	return errors.New("postgres storage does not support function SetUnitProperties") // TODO: implement me
}
