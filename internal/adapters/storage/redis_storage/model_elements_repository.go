package redis_storage

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type ModelElementsRedisRepository struct {
	client  *redis.Client
	keyType string
}

var modelElementID int64 = 0

func NewModelElementsRedisRepository(host, port, password, keyType string) *ModelElementsRedisRepository {
	return &ModelElementsRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func (r *ModelElementsRedisRepository) GetAllModelElements(conn storage.ConnDB, layer string) ([]entities.ModelElement, error) {
	keyPattern := compositeKey(r.keyType, layer, "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	modelElements := make([]entities.ModelElement, 0, len(keys))
	keyFormat := compositeKey(r.keyType, layer, "%d")
	for _, key := range keys {
		var elem entities.ModelElement
		_, err = fmt.Sscanf(key, keyFormat, &elem.ID)
		if err != nil {
			return nil, err
		}
		elem.Name, err = r.client.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		modelElements = append(modelElements, elem)
	}

	return modelElements, nil
}

func (r *ModelElementsRedisRepository) SaveModelElements(conn storage.ConnDB, layer string, modelElements []string) ([]int, error) { // TODO: fix full duplicates (exclude id)
	idArray := make([]int, 0)
	var err error
	for _, elem := range modelElements {
		idStr := strconv.FormatInt(modelElementID, 10)
		key := compositeKey(r.keyType, layer, idStr)
		err = r.client.Set(context.Background(), key, elem, 0).Err()
		if err != nil {
			break
		}
		idArray = append(idArray, int(modelElementID))
		modelElementID++
	}

	if err != nil {
		keyPrefix := compositeKey(r.keyType, layer)
		_ = delKeysByID(r.client, keyPrefix, idArray)
		return nil, err
	}

	return idArray, nil
}
