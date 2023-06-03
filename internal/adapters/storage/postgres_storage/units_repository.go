package postgres_storage

import (
	"context"
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/sqlx_utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sync"
)

type UnitsPgRepository struct{}

func (r *UnitsPgRepository) makeOutputUnitDTO(conn storage.ConnDB, layer, lang string, unit entities.Unit) (interfaces.OutputUnitDTO, error) {
	args := map[string]any{
		"layer_name": layer,
		"lang":       lang,
		"unit_id":    unit.ID,
	}
	propertiesID, err := namedSelectSliceFromScript[[]int](conn, selectPropertiesIdByUnitId, args)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	contextsID, err := namedSelectSliceFromScript[[]int](conn, selectContextsIdByUnit, args)
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
			unitDTO, err := r.makeOutputUnitDTO(conn, layer, lang, unit)
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
			unitDTO, err := r.makeOutputUnitDTO(conn, layer, lang, unit)
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

type unitQueries struct {
	unlinkedUnitsQuery string
	linkedUnitsQuery   string
	optionalArgs       map[string]any
}

func (r *UnitsPgRepository) getUnitsByQueries(conn storage.ConnDB, layer string, langs []string, qData unitQueries) (joinedUnitMaps, entities.UnitsMap, error) {
	args1 := map[string]any{
		"layer_name": layer,
		"lang":       "",
	}
	args2 := map[string]any{
		"layer_name": layer,
	}
	for argName, arg := range qData.optionalArgs {
		args1[argName] = arg
		args2[argName] = arg
	}

	unlinkedUnitsMap := make(entities.UnitsMap)
	for _, lang := range langs {
		args1["lang"] = lang
		unlinkedUnits, err := namedSelectSliceFromScript[[]entities.Unit](conn, qData.unlinkedUnitsQuery, args1)
		if err != nil {
			return nil, nil, err
		}
		unlinkedUnitsMap[lang] = unlinkedUnits
	}

	linkedUnits, err := namedSelectSliceFromScript[joinedUnits](conn, qData.linkedUnitsQuery, args2)
	if err != nil {
		return nil, nil, err
	}

	return linkedUnits.toMaps(), unlinkedUnitsMap, nil
}

func (r *UnitsPgRepository) GetAllUnits(conn storage.ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	linkedUnits, unlinkedUnits, err := r.getUnitsByQueries(conn, layer, []string{"ru", "en"}, unitQueries{
		unlinkedUnitsQuery: selectUnlinkedUnits,
		linkedUnitsQuery:   selectAllLinkedUnits,
		optionalArgs:       nil,
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	combinedUnits, contextsID, err := r.combineUnits(conn, layer, linkedUnits, unlinkedUnits)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	contexts, err := selectSliceFromScript[[]interfaces.ContextDTO](conn, selectContextsById, pq.Array(contextsID))
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{Units: combinedUnits, Contexts: contexts}, nil
}

func (r *UnitsPgRepository) GetUnitsByModels(conn storage.ConnDB, layer string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	linkedUnits, unlinkedUnits, err := r.getUnitsByQueries(conn, layer, []string{"ru", "en"}, unitQueries{
		unlinkedUnitsQuery: selectUnlinkedUnitsAndModelsId,
		linkedUnitsQuery:   selectLinkedUnitsByModelsId,
		optionalArgs: map[string]any{
			"models_id_array": pq.Array(modelsID),
		},
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	combinedUnits, contextsID, err := r.combineUnits(conn, layer, linkedUnits, unlinkedUnits)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	contexts, err := selectSliceFromScript[[]interfaces.ContextDTO](conn, selectContextsById, pq.Array(contextsID))
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{Units: combinedUnits, Contexts: contexts}, nil
}

func (r *UnitsPgRepository) GetUnitsByProperties(conn storage.ConnDB, layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	linkedUnits, unlinkedUnits, err := r.getUnitsByQueries(conn, layer, []string{"ru", "en"}, unitQueries{
		unlinkedUnitsQuery: selectUnlinkedUnitsAndPropertiesId,
		linkedUnitsQuery:   selectLinkedUnitsByPropertiesId,
		optionalArgs: map[string]any{
			"properties_id_array": pq.Array(propertiesID),
		},
	})
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	combinedUnits, contextsID, err := r.combineUnits(conn, layer, linkedUnits, unlinkedUnits)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	contexts, err := selectSliceFromScript[[]interfaces.ContextDTO](conn, selectContextsById, pq.Array(contextsID))
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{Units: combinedUnits, Contexts: contexts}, nil
}

func (r *UnitsPgRepository) insertContext(tx sqlx.ExtContext, ctxText string) (int, error) {
	return selectValueFromScript[int](tx, insertContext, ctxText)
}

func (r *UnitsPgRepository) insertUnits(tx sqlx.ExtContext, layer, lang string, modelsID []int, unitTexts []string) ([]int, error) {
	if len(modelsID) != len(unitTexts) {
		return nil, errors.New("incorrect unit DTO parsing")
	}
	args := map[string]any{
		"layer_name": layer,
		"lang":       lang,
		"models_id":  pq.Array(modelsID),
		"unit_texts": pq.Array(unitTexts),
	}
	return namedSelectSliceFromScript[[]int](tx, insertUnits, args)
}

func (r *UnitsPgRepository) insertUnitProperties(tx sqlx.ExtContext, layer, lang string, unitID int, propertiesID []int) error {
	args := map[string]any{
		"layer_name":    layer,
		"lang":          lang,
		"unit_id":       unitID,
		"properties_id": pq.Array(propertiesID),
	}
	return executeNamedScript(tx, insertUnitProperties, args)
}

func (r *UnitsPgRepository) insertContextUnits(tx sqlx.ExtContext, layer, lang string, contextID int, unitsID []int) error {
	args := map[string]any{
		"layer_name": layer,
		"lang":       lang,
		"context_id": contextID,
		"units_id":   pq.Array(unitsID),
	}
	return executeNamedScript(tx, insertContextUnits, args)
}

func (r *UnitsPgRepository) linkUnits(tx sqlx.ExtContext, layer, unitRu, unitEn string) error {
	args := map[string]any{
		"layer_name": layer,
		"unit_ru":    unitRu,
		"unit_en":    unitEn,
	}
	return executeNamedScript(tx, linkUnits, args)
}

func (r *UnitsPgRepository) saveUnitsTX(ctx context.Context, tx sqlx.ExtContext, layer string, data interfaces.SaveUnitsDTO) error {
	wg := sync.WaitGroup{}
	var globalErr error

	for lang, ctxText := range data.Contexts {
		contextID, err := r.insertContext(tx, ctxText)
		if err != nil {
			return err
		}

		modelsID := make([]int, 0, len(data.Units))
		unitTexts := make([]string, 0, len(data.Units))
		for _, linkedUnitsDTO := range data.Units {
			unitDTO, inMap := linkedUnitsDTO[lang]
			if !inMap {
				continue
			}
			modelsID = append(modelsID, unitDTO.ModelID)
			unitTexts = append(unitTexts, unitDTO.Text)
		}

		unitsID, err := r.insertUnits(tx, layer, lang, modelsID, unitTexts)
		if err != nil {
			return err
		}

		go func(lang string, unitsID []int) {
			wg.Add(1)
			defer wg.Done()
			for i, linkedUnitsDTO := range data.Units {
				unitDTO, inMap := linkedUnitsDTO[lang]
				if !inMap || len(unitDTO.PropertiesID) == 0 {
					continue
				}
				err = r.insertUnitProperties(tx, layer, lang, unitsID[i], unitDTO.PropertiesID)
				if err != nil {
					globalErr = err
				}
			}
		}(lang, unitsID)

		go func(lang string, contextID int, unitsID []int) {
			wg.Add(1)
			defer wg.Done()
			err = r.insertContextUnits(tx, layer, lang, contextID, unitsID)
			if err != nil {
				globalErr = err
			}
		}(lang, contextID, unitsID)

	}

	for _, linkedUnitsDTO := range data.Units {
		unitRuDTO, inMap := linkedUnitsDTO["ru"]
		if !inMap {
			continue
		}
		unitEnDTO, inMap := linkedUnitsDTO["en"]
		if !inMap {
			continue
		}
		err := r.linkUnits(tx, layer, unitRuDTO.Text, unitEnDTO.Text)
		if err != nil {
			globalErr = err
		}
	}

	wg.Wait()

	return globalErr
}

func (r *UnitsPgRepository) SaveUnits(conn storage.ConnDB, layer string, data interfaces.SaveUnitsDTO) error {
	sqlxDB, ok := conn.(sqlx_utils.TxRunner)
	if !ok {
		return errors.New("cannot get TxRunner from argument")
	}

	return sqlx_utils.RunTx(context.Background(), sqlxDB, func(tx *sqlx.Tx) error {
		err := r.saveUnitsTX(context.Background(), tx, layer, data)
		return err
	})
}

func (r *UnitsPgRepository) RenameUnit(conn storage.ConnDB, layer, lang, oldName, newName string) error {
	args := map[string]any{
		"layer_name": layer,
		"lang":       lang,
		"old_name":   oldName,
		"new_name":   newName,
	}
	return executeNamedScript(conn, updateUnitNames, args)
}

func (r *UnitsPgRepository) setUnitPropertiesTX(ctx context.Context, tx sqlx.ExtContext, layer, lang, unitName string, propertiesID []int) error {
	args := map[string]any{
		"layer_name":    layer,
		"lang":          lang,
		"unit_name":     unitName,
		"properties_id": pq.Array(propertiesID),
	}
	return executeNamedScript(tx, updateUnitProperties, args)
}

func (r *UnitsPgRepository) SetUnitProperties(conn storage.ConnDB, layer, lang, unitName string, propertiesID []int) error {
	sqlxDB, ok := conn.(sqlx_utils.TxRunner)
	if !ok {
		return errors.New("cannot get TxRunner from argument")
	}

	return sqlx_utils.RunTx(context.Background(), sqlxDB, func(tx *sqlx.Tx) error {
		err := r.setUnitPropertiesTX(context.Background(), tx, layer, lang, unitName, propertiesID)
		return err
	})
}
