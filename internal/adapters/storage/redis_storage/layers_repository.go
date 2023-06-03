package redis_storage

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
)

type LayersRedisRepository struct {
	client  *redis.Client
	keyType string
}

func NewLayersRedisRepository(host, port, password, keyType string) *LayersRedisRepository {
	return &LayersRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func (r *LayersRedisRepository) LayerExist(conn storage.ConnDB, layer string) (bool, error) {
	key := compositeKey(r.keyType, layer)
	val, err := r.client.Get(context.Background(), key).Result()
	switch {
	case err == redis.Nil:
		return false, nil
	case err != nil:
		return false, err
	case val == existFlag:
		return true, nil
	default:
		return false, nil
	}
}

func (r *LayersRedisRepository) GetAllLayers(conn storage.ConnDB) ([]string, error) {
	keyPattern := compositeKey(r.keyType, "*")
	keys, err := getKeysByPattern(r.client, keyPattern)
	if err != nil {
		return nil, err
	}

	layers := make([]string, 0, len(keys))
	keyFormat := compositeKey(r.keyType, "%s")
	for _, key := range keys {
		var layer string
		_, err = fmt.Sscanf(key, keyFormat, &layer)
		if err != nil {
			return nil, err
		}
		layers = append(layers, layer)
	}

	return layers, nil
}

func (r *LayersRedisRepository) SaveLayer(conn storage.ConnDB, name string) error {
	key := compositeKey(r.keyType, name)
	return r.client.Set(context.Background(), key, existFlag, 0).Err()
}
