package postgres_storage

import "github.com/Inspirate789/Thermy-backend/internal/domain/entities"

type joinedUnit struct {
	UnitRuID      int    `db:"unit_ru_id"`
	UnitRuModelID int    `db:"unit_ru_model_id"`
	UnitRuRegDate string `db:"unit_ru_registration_date"`
	UnitRuText    string `db:"unit_ru_text"`
	UnitEnID      int    `db:"unit_en_id"`
	UnitEnModelID int    `db:"unit_en_model_id"`
	UnitEnRegDate string `db:"unit_en_registration_date"`
	UnitEnText    string `db:"unit_en_text"`
}

type joinedUnits []joinedUnit

type joinedUnitMap map[string]entities.Unit

type joinedUnitMaps []joinedUnitMap

func (ju *joinedUnit) toMap() joinedUnitMap {
	return map[string]entities.Unit{
		"ru": {
			ID:      ju.UnitRuID,
			ModelID: ju.UnitRuModelID,
			RegDate: ju.UnitRuRegDate,
			Text:    ju.UnitRuText,
		},
		"en": {
			ID:      ju.UnitEnID,
			ModelID: ju.UnitEnModelID,
			RegDate: ju.UnitEnRegDate,
			Text:    ju.UnitEnText,
		},
	}
}

func (ju joinedUnits) toMaps() joinedUnitMaps {
	unitMaps := make(joinedUnitMaps, 0, len(ju))
	for _, unit := range ju {
		unitMaps = append(unitMaps, unit.toMap())
	}

	return unitMaps
}
