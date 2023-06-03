package redis_storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
)

type UnitsRedisRepository struct {
	client  *redis.Client
	keyType string
}

type unit struct {
	ModelID int    `json:"model_id"`
	Text    string `json:"text"`
}

const (
	unitPrefix         = "units"
	userLinkPrefix     = "users"
	propertyLinkPrefix = "properties"
	contextPrefix      = "contexts"
	contextLinkPrefix  = "contextLinks"
	unitLinkPrefix     = "link"
)

var (
	unitID    int64 = 0
	contextID int64 = 0
)

func NewUnitsRedisRepository(host, port, password, keyType string) *UnitsRedisRepository {
	return &UnitsRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func compositeUnitKeyPrefix(keyType, layer, lang string) string {
	return compositeKey(keyType, unitPrefix, layer, lang)
}

func compositeUnitKey(keyType, layer, lang string, id int64) string {
	idStr := strconv.FormatInt(id, 10)
	return compositeKey(compositeUnitKeyPrefix(keyType, layer, lang), idStr)
}

func compositeUserLinkKey(keyType, layer, lang string, id int64) string {
	idStr := strconv.FormatInt(id, 10)
	return compositeKey(keyType, userLinkPrefix, layer, lang, idStr)
}

func compositePropertyLinkKey(keyType, layer, lang string, id int64) string {
	idStr := strconv.FormatInt(id, 10)
	return compositeKey(keyType, propertyLinkPrefix, layer, lang, idStr)
}

func (r *UnitsRedisRepository) getUnitPropertiesID(layer, lang string, unitID int) ([]int, error) { // TODO: fix
	key := compositePropertyLinkKey(r.keyType, layer, lang, int64(unitID))
	data, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var propertiesID []int
	err = json.Unmarshal([]byte(data), &propertiesID)
	if err != nil {
		return nil, err
	}

	return propertiesID, nil
}

func (r *UnitsRedisRepository) getUnitContextsID(layer, lang string, unitID int) ([]int, error) {
	unitIdStr := strconv.FormatInt(int64(unitID), 10)
	keyPattern := compositeKey(r.keyType, contextLinkPrefix, layer, lang, "*", unitIdStr)
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	contextsID := make([]int, 0, len(keys))
	keyFormat := compositeKey(r.keyType, contextLinkPrefix, layer, lang, "%d", unitIdStr)
	for _, key := range keys {
		var id int
		_, err = fmt.Sscanf(key, keyFormat, &id)
		if err != nil {
			return nil, err
		}
		contextsID = append(contextsID, id)
	}

	return contextsID, nil
}

func (r *UnitsRedisRepository) makeOutputUnitDTO(layer, lang string, unit entities.Unit) (interfaces.OutputUnitDTO, error) {
	propertiesID, err := r.getUnitPropertiesID(layer, lang, unit.ID)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	contextsID, err := r.getUnitContextsID(layer, lang, unit.ID)
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

func (r *UnitsRedisRepository) getUnit(key, keyFormat, layer, lang string) (interfaces.OutputUnitDTO, error) {
	var curUnit entities.Unit
	_, err := fmt.Sscanf(key, keyFormat, &curUnit.ID)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	data, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	var unitData unit
	err = json.Unmarshal([]byte(data), &unitData)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}
	curUnit.ModelID = unitData.ModelID
	curUnit.Text = unitData.Text

	unitDTO, err := r.makeOutputUnitDTO(layer, lang, curUnit)
	if err != nil {
		return interfaces.OutputUnitDTO{}, err
	}

	return unitDTO, nil
}

func (r *UnitsRedisRepository) getUnits(layer, lang string) ([]interfaces.OutputUnitDTO, error) {
	keyPattern := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	unitsDTO := make([]interfaces.OutputUnitDTO, 0, len(keys)) // TODO: work with registration date
	keyFormat := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "%d")
	for _, key := range keys {
		unitDTO, err := r.getUnit(key, keyFormat, layer, lang)
		if err != nil {
			return nil, err
		}

		unitsDTO = append(unitsDTO, unitDTO)
	}

	return unitsDTO, nil
}

func (r *UnitsRedisRepository) getUnitsByModels(layer, lang string, modelsID []int) ([]interfaces.OutputUnitDTO, error) {
	keyPattern := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	unitsDTO := make([]interfaces.OutputUnitDTO, 0, len(keys)) // TODO: work with registration date
	keyFormat := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "%d")
	for _, key := range keys {
		unitDTO, err := r.getUnit(key, keyFormat, layer, lang)
		if err != nil {
			return nil, err
		}

		if containElem[[]int](unitDTO.ModelID, modelsID) {
			unitsDTO = append(unitsDTO, unitDTO)
		}
	}

	return unitsDTO, nil
}

func (r *UnitsRedisRepository) getUnitsByProperties(layer, lang string, propertiesID []int) ([]interfaces.OutputUnitDTO, error) {
	keyPattern := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	unitsDTO := make([]interfaces.OutputUnitDTO, 0, len(keys)) // TODO: work with registration date
	keyFormat := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "%d")
	for _, key := range keys {
		unitDTO, err := r.getUnit(key, keyFormat, layer, lang)
		if err != nil {
			return nil, err
		}

		if containSlice[[]int](unitDTO.PropertiesID, propertiesID) {
			unitsDTO = append(unitsDTO, unitDTO)
		}
	}

	return unitsDTO, nil
}

func makeUnitsSet(units []interfaces.OutputUnitDTO) map[string]interfaces.OutputUnitDTO {
	set := make(map[string]interfaces.OutputUnitDTO)
	for _, curUnit := range units {
		set[curUnit.Text] = curUnit
	}

	return set
}

func (r *UnitsRedisRepository) getUnitTranslatesSub(layer, varKeyPattern, varKeyFormat string) ([]string, error) {
	keyPattern := compositeKey(r.keyType, unitLinkPrefix, layer, varKeyPattern)

	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	translates := make([]string, 0, len(keys))
	keyFormat := compositeKey(r.keyType, unitLinkPrefix, layer, varKeyFormat)
	for _, key := range keys {
		var translate string
		_, err = fmt.Sscanf(key, keyFormat, &translate)
		if err != nil {
			return nil, err
		}
		translates = append(translates, translate)
	}

	return removeDuplicates[[]string](translates), nil
}

func (r *UnitsRedisRepository) getUnitTranslates(layer, unitText string) ([]string, error) {
	varKeyPattern := compositeKey(unitText, "*")
	varKeyFormat := compositeKey(unitText, "%s")
	translates1, err := r.getUnitTranslatesSub(layer, varKeyPattern, varKeyFormat)
	if err != nil {
		return nil, err
	}

	varKeyPattern = compositeKey("*", unitText)
	varKeyFormat = compositeKey("%s", unitText)
	translates2, err := r.getUnitTranslatesSub(layer, varKeyPattern, varKeyFormat)
	if err != nil {
		return nil, err
	}

	return append(translates1, translates2...), nil
}

func (r *UnitsRedisRepository) extractContextsID(units []interfaces.OutputUnitDTO) []int {
	contextsID := make([]int, 0)
	for _, curUnit := range units {
		contextsID = append(contextsID, curUnit.ContextsID...)
	}

	return removeDuplicates[[]int](contextsID)
}

func (r *UnitsRedisRepository) extractContexts(units []interfaces.OutputUnitDTO) ([]interfaces.ContextDTO, error) {
	contextsID := r.extractContextsID(units)
	contexts := make([]interfaces.ContextDTO, 0, len(contextsID))
	for _, id := range contextsID {
		ctx := interfaces.ContextDTO{ID: id}
		idStr := strconv.FormatInt(int64(id), 10)
		key := compositeKey(r.keyType, contextPrefix, idStr)
		var err error
		ctx.Text, err = r.client.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		contexts = append(contexts, ctx)
	}

	return contexts, nil
}

func (r *UnitsRedisRepository) makeOutputUnitsDTO(layer string, unitsRu, unitsEn []interfaces.OutputUnitDTO) (interfaces.OutputUnitsDTO, error) { // TODO: fix code style
	setUnitsRu := makeUnitsSet(unitsRu)
	setUnitsEn := makeUnitsSet(unitsEn)

	units := make(interfaces.UnitDtoMaps, 0)
	addedUnitsRu := make(map[string]bool)
	addedUnitsEn := make(map[string]bool)
	for textRu, unitRu := range setUnitsRu {
		_, exist := addedUnitsRu[textRu]
		if exist {
			continue
		}

		translates, err := r.getUnitTranslates(layer, textRu)
		if err != nil {
			return interfaces.OutputUnitsDTO{}, err
		}

		if len(translates) == 0 {
			units = append(units, map[string]interfaces.OutputUnitDTO{"ru": unitRu})
		}

		for _, textEn := range translates {
			_, exist = addedUnitsEn[textEn]
			if exist {
				continue
			}
			units = append(units, map[string]interfaces.OutputUnitDTO{
				"ru": unitRu,
				"en": setUnitsEn[textEn],
			})
			addedUnitsEn[textEn] = true
		}

		addedUnitsRu[textRu] = true
	}
	for textEn, unitEn := range setUnitsEn {
		_, exist := addedUnitsEn[textEn]
		if exist {
			continue
		}
		units = append(units, map[string]interfaces.OutputUnitDTO{"en": unitEn})
		addedUnitsEn[textEn] = true
	}

	contexts, err := r.extractContexts(append(unitsRu, unitsEn...))
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return interfaces.OutputUnitsDTO{
		Units:    units,
		Contexts: contexts,
	}, nil
}

func (r *UnitsRedisRepository) GetAllUnits(_ storage.ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	unitsRu, err := r.getUnits(layer, "ru")
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}
	unitsEn, err := r.getUnits(layer, "en")
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return r.makeOutputUnitsDTO(layer, unitsRu, unitsEn)
}

func (r *UnitsRedisRepository) GetUnitsByModels(_ storage.ConnDB, layer string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	unitsRu, err := r.getUnitsByModels(layer, "ru", modelsID)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}
	unitsEn, err := r.getUnitsByModels(layer, "en", modelsID)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return r.makeOutputUnitsDTO(layer, unitsRu, unitsEn)
}

func (r *UnitsRedisRepository) GetUnitsByProperties(_ storage.ConnDB, layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	unitsRu, err := r.getUnitsByProperties(layer, "ru", propertiesID)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}
	unitsEn, err := r.getUnitsByProperties(layer, "en", propertiesID)
	if err != nil {
		return interfaces.OutputUnitsDTO{}, err
	}

	return r.makeOutputUnitsDTO(layer, unitsRu, unitsEn)
}

func (r *UnitsRedisRepository) insertContext(ctxText string) (int, error) {
	idStr := strconv.FormatInt(contextID, 10)
	key := compositeKey(r.keyType, contextPrefix, idStr)
	err := r.client.Set(context.Background(), key, ctxText, 0).Err()
	if err != nil {
		return 0, err
	}

	id := contextID
	contextID++

	return int(id), nil
}

func (r *UnitsRedisRepository) extractUnitData(data interfaces.SaveUnitsDTO, lang string) ([]int, []string) {
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

	return modelsID, unitTexts
}

func (r *UnitsRedisRepository) insertUnits(layer, lang, username string, modelsID []int, unitTexts []string) ([]int, error) {
	if len(modelsID) != len(unitTexts) {
		return nil, errors.New("incorrect unit DTO parsing")
	}

	idArray := make([]int, 0)
	var globalErr error
	for i := range unitTexts {
		curUnit := unit{
			ModelID: modelsID[i],
			Text:    unitTexts[i],
		}
		data, err := json.Marshal(curUnit)
		if err != nil {
			globalErr = err
			break
		}

		key := compositeUnitKey(r.keyType, layer, lang, unitID)
		err = r.client.Set(context.Background(), key, data, 0).Err() // TODO: fix duplicates?
		if err != nil {
			globalErr = err
			break
		}

		key = compositeUserLinkKey(r.keyType, layer, lang, unitID)
		err = r.client.Set(context.Background(), key, username, 0).Err()
		if err != nil {
			globalErr = err
			break
		}

		idArray = append(idArray, int(unitID))
		unitID++
	}

	if globalErr != nil {
		keyPrefix := compositeUnitKeyPrefix(r.keyType, layer, lang)
		_ = delKeysByID(r.client, keyPrefix, idArray)
		return nil, globalErr
	}

	return idArray, nil
}

func (r *UnitsRedisRepository) insertUnitProperties(layer, lang string, unitID int, propertiesID []int) error {
	data, err := json.Marshal(propertiesID)
	if err != nil {
		return err
	}

	key := compositePropertyLinkKey(r.keyType, layer, lang, int64(unitID))

	return r.client.Set(context.Background(), key, data, 0).Err()
}

func (r *UnitsRedisRepository) saveUnitProperties(units []map[string]interfaces.SaveUnitDTO, unitsID []int, layer, lang string) error {
	for i, linkedUnitsDTO := range units {
		unitDTO, inMap := linkedUnitsDTO[lang]
		if !inMap || len(unitDTO.PropertiesID) == 0 {
			continue
		}
		err := r.insertUnitProperties(layer, lang, unitsID[i], unitDTO.PropertiesID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UnitsRedisRepository) insertContextUnits(layer, lang string, contextID int, unitsID []int) error {
	contextIdStr := strconv.FormatInt(int64(contextID), 10)
	var err error
	for _, id := range unitsID {
		unitIdStr := strconv.FormatInt(int64(id), 10)
		key := compositeKey(r.keyType, contextLinkPrefix, layer, lang, contextIdStr, unitIdStr)
		err = r.client.Set(context.Background(), key, existFlag, 0).Err()
		if err != nil {
			break
		}
	}

	if err != nil {
		keyPrefix := compositeKey(r.keyType, contextLinkPrefix, layer, lang, contextIdStr)
		_ = delKeysByID(r.client, keyPrefix, unitsID) // TODO: delete only existing keys
		return err
	}

	return nil
}

func (r *UnitsRedisRepository) linkUnitPair(layer, unitRu, unitEn string) error {
	key := compositeKey(r.keyType, unitLinkPrefix, layer, unitRu, unitEn)
	return r.client.Set(context.Background(), key, existFlag, 0).Err()
}

func (r *UnitsRedisRepository) linkUnits(units []map[string]interfaces.SaveUnitDTO, layer string) error {
	for _, linkedUnitsDTO := range units {
		unitRuDTO, inMap := linkedUnitsDTO["ru"]
		if !inMap {
			continue
		}
		unitEnDTO, inMap := linkedUnitsDTO["en"]
		if !inMap {
			continue
		}
		err := r.linkUnitPair(layer, unitRuDTO.Text, unitEnDTO.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UnitsRedisRepository) SaveUnits(conn storage.ConnDB, layer string, data interfaces.SaveUnitsDTO) error {
	username, ok := conn.(string)
	if !ok {
		return errors.New("cannot get string from argument")
	}

	wg := sync.WaitGroup{}
	var globalErr error
	for lang, ctxText := range data.Contexts {
		ctxID, err := r.insertContext(ctxText)
		if err != nil {
			return err
		}

		modelsID, unitTexts := r.extractUnitData(data, lang)
		unitsID, err := r.insertUnits(layer, lang, username, modelsID, unitTexts)
		if err != nil {
			return err
		}

		go func(lang string, unitsID []int) {
			wg.Add(1)
			defer wg.Done()
			err = r.saveUnitProperties(data.Units, unitsID, layer, lang)
			if err != nil {
				globalErr = err
			}
		}(lang, unitsID)

		go func(lang string, contextID int, unitsID []int) {
			wg.Add(1)
			defer wg.Done()
			err = r.insertContextUnits(layer, lang, contextID, unitsID)
			if err != nil {
				globalErr = err
			}
		}(lang, ctxID, unitsID)
	}

	err := r.linkUnits(data.Units, layer)
	if err != nil {
		globalErr = err
	}

	wg.Wait()

	return globalErr
}

func (r *UnitsRedisRepository) RenameUnit(_ storage.ConnDB, layer, lang, oldName, newName string) error {
	keyPattern := compositeKey(compositeUnitKeyPrefix(r.keyType, layer, lang), "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return err
	}

	for _, key := range keys {
		name, err := r.client.Get(context.Background(), key).Result()
		if err == nil && name == oldName {
			err = r.client.Set(context.Background(), key, newName, 0).Err()
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func getUnitIdByName(client *redis.Client, keyType, layer, lang, name string) (int, error) {
	keyPrefix := compositeUnitKeyPrefix(keyType, layer, lang)
	keyPattern := compositeKey(keyPrefix, "*")
	keys, err := getKeysByPattern(client, keyPattern)
	if err != nil {
		return 0, err
	}

	keyFormat := compositeKey(keyPrefix, "%d")
	for _, key := range keys {
		val, err := client.Get(context.Background(), key).Result()
		if err == nil && val == name {
			var id int
			_, err = fmt.Sscanf(key, keyFormat, &id)
			if err != nil {
				return 0, err
			}
			return id, nil
		}
	}

	return 0, fmt.Errorf("cannot find unit with name %s", name)
}

func (r *UnitsRedisRepository) SetUnitProperties(_ storage.ConnDB, layer, lang, unitName string, propertiesID []int) error {
	id, err := getUnitIdByName(r.client, r.keyType, layer, lang, unitName)
	if err != nil {
		return err
	}
	return r.insertUnitProperties(layer, lang, id, propertiesID)
}
