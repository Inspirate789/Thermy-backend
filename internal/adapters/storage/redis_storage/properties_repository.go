package redis_storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type PropertiesRedisRepository struct {
	client  *redis.Client
	keyType string
}

var propertyID int64 = 0

func NewPropertiesRedisRepository(host, port, password, keyType string) *PropertiesRedisRepository {
	return &PropertiesRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func (r *PropertiesRedisRepository) GetAllProperties(_ storage.ConnDB) ([]entities.Property, error) {
	keyPattern := compositeKey(r.keyType, "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	properties := make([]entities.Property, 0, len(keys))
	keyFormat := compositeKey(r.keyType, "%d")
	for _, key := range keys {
		var property entities.Property
		_, err = fmt.Sscanf(key, keyFormat, &property.ID)
		if err != nil {
			return nil, err
		}
		property.Name, err = r.client.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func (r *PropertiesRedisRepository) getPropertiesByID(propertiesID []int) ([]entities.Property, error) {
	properties := make([]entities.Property, 0, len(propertiesID))
	for _, id := range propertiesID {
		property := entities.Property{ID: id}
		idStr := strconv.FormatInt(int64(id), 10)
		key := compositeKey(r.keyType, idStr)
		var err error
		property.Name, err = r.client.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func (r *PropertiesRedisRepository) GetPropertiesByUnit(_ storage.ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	id, err := getUnitIdByName(r.client, r.keyType, layer, unit.Lang, unit.Text)
	if err != nil {
		return nil, err
	}

	idStr := strconv.FormatInt(int64(id), 10)
	key := compositeKey(r.keyType, propertyLinkPrefix, layer, unit.Lang, idStr)

	data, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var propertiesID []int
	err = json.Unmarshal([]byte(data), &propertiesID)
	if err != nil {
		return nil, err
	}

	return r.getPropertiesByID(propertiesID)
}

func (r *PropertiesRedisRepository) SaveProperties(_ storage.ConnDB, properties []string) ([]int, error) {
	idArray := make([]int, 0)
	var err error
	for _, elem := range properties {
		idStr := strconv.FormatInt(propertyID, 10)
		key := compositeKey(r.keyType, idStr)
		err = r.client.Set(context.Background(), key, elem, 0).Err()
		if err != nil {
			break
		}
		idArray = append(idArray, int(propertyID))
		propertyID++
	}

	if err != nil {
		keyPrefix := compositeKey(r.keyType)
		_ = delKeysByID(r.client, keyPrefix, idArray)
		return nil, err
	}

	return idArray, nil
}
