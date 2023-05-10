package redis_storage

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type ModelsRedisRepository struct {
	client  *redis.Client
	keyType string
}

var modelID int64 = 0

func NewModelsRedisRepository(host, port, password, keyType string) *ModelsRedisRepository {
	return &ModelsRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func (r *ModelsRedisRepository) GetAllModels(conn storage.ConnDB, layer string) ([]entities.Model, error) {
	keyPattern := compositeKey(r.keyType, layer, "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	models := make([]entities.Model, 0, len(keys))
	keyFormat := compositeKey(r.keyType, layer, "%d")
	for _, key := range keys {
		var elem entities.Model
		_, err = fmt.Sscanf(key, keyFormat, &elem.ID)
		if err != nil {
			return nil, err
		}
		elem.Name, err = r.client.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		models = append(models, elem)
	}

	return models, nil
}

func (r *ModelsRedisRepository) SaveModels(conn storage.ConnDB, layer string, models []string) ([]int, error) {
	idArray := make([]int, 0)
	var err error
	for _, elem := range models {
		idStr := strconv.FormatInt(modelID, 10)
		key := compositeKey(r.keyType, layer, idStr)
		err = r.client.Set(context.Background(), key, elem, 0).Err()
		if err != nil {
			break
		}
		idArray = append(idArray, int(modelID))
		modelID++
	}

	if err != nil {
		keyPrefix := compositeKey(r.keyType, layer)
		_ = delKeysByID(r.client, keyPrefix, idArray)
		return nil, err
	}

	return idArray, nil
}
