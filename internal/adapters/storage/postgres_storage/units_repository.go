package postgres_storage

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/lib/pq"
)

type UnitsPgRepository struct{}

func makeOutputUnitDTO(conn storage.ConnDB, layer string, lang string, unit entities.Unit) (interfaces.OutputUnitDTO, error) {
	args := map[string]any{
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

func (r *UnitsPgRepository) combineUnlinkedUnits(conn storage.ConnDB, layer string, unlinkedUnits entities.UnitsMap, combinedUnits interfaces.UnitDtoMaps) (interfaces.UnitDtoMaps, []int, error) {
	uniqueContextsID := make(map[int]bool)

	for lang := range unlinkedUnits {
		for _, unit := range unlinkedUnits[lang] {
			unitDTO, err := makeOutputUnitDTO(conn, layer, lang, unit)
			if err != nil {
				return nil, nil, err
			}

			for _, contextID := range unitDTO.ContextsID {
				uniqueContextsID[contextID] = true
			}

			combinedUnits = append(combinedUnits, map[string]interfaces.OutputUnitDTO{
				lang: unitDTO,
			})
		}
	}

	contextsID := make([]int, 0, len(uniqueContextsID))
	for id := range uniqueContextsID {
		contextsID = append(contextsID, id)
	}

	return combinedUnits, contextsID, nil
}

func (r *UnitsPgRepository) combineLinkedUnits(conn storage.ConnDB, layer string, linkedUnits joinedUnitMaps, combinedUnits interfaces.UnitDtoMaps) (interfaces.UnitDtoMaps, []int, error) {
	uniqueContextsID := make(map[int]bool)

	for _, unitMap := range linkedUnits {
		unitMapDTO := make(map[string]interfaces.OutputUnitDTO)
		for lang := range unitMap {
			unit := unitMap[lang]
			unitDTO, err := makeOutputUnitDTO(conn, layer, lang, unit)
			if err != nil {
				return nil, nil, err
			}
			unitMapDTO[lang] = unitDTO

			for _, contextID := range unitDTO.ContextsID {
				uniqueContextsID[contextID] = true
			}
		}

		combinedUnits = append(combinedUnits, unitMapDTO)
	}

	contextsID := make([]int, 0, len(uniqueContextsID))
	for id := range uniqueContextsID {
		contextsID = append(contextsID, id)
	}

	return combinedUnits, contextsID, nil
}

func (r *UnitsPgRepository) combineUnits(conn storage.ConnDB, layer string, linkedUnits joinedUnitMaps, unlinkedUnits entities.UnitsMap) (interfaces.UnitDtoMaps, []int, error) {
	unlinkedUnitsLen := 0
	for lang := range unlinkedUnits {
		unlinkedUnitsLen += len(unlinkedUnits[lang])
	}

	combinedUnits := make(interfaces.UnitDtoMaps, 0, unlinkedUnitsLen+len(linkedUnits))

	combinedUnits, contextsID1, err := r.combineUnlinkedUnits(conn, layer, unlinkedUnits, combinedUnits)
	if err != nil {
		return nil, nil, err
	}

	combinedUnits, contextsID2, err := r.combineLinkedUnits(conn, layer, linkedUnits, combinedUnits)
	if err != nil {
		return nil, nil, err
	}

	return combinedUnits, append(contextsID1, contextsID2...), nil
}

func (r *UnitsPgRepository) GetAllUnits(conn storage.ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	unlinkedUnitsRu, err := namedSelectSliceFromScript[[]entities.Unit](conn, selectUnlinkedUnitsByLangQuery, map[string]any{
		"layer_name": layer,
		"lang":       "ru",
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	unlinkedUnitsEn, err := namedSelectSliceFromScript[[]entities.Unit](conn, selectUnlinkedUnitsByLangQuery, map[string]any{
		"layer_name": layer,
		"lang":       "en",
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	linkedUnits, err := namedSelectSliceFromScript[joinedUnits](conn, selectAllLinkedUnitsQuery, map[string]any{
		"layer_name": layer,
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	combinedUnits, contextsID, err := r.combineUnits(conn, layer, linkedUnits.toMaps(), entities.UnitsMap{
		"ru": unlinkedUnitsRu,
		"en": unlinkedUnitsEn,
	})

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
