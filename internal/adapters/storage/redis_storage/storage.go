package redis_storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	*LayersRedisRepository
	*ModelsRedisRepository
	*ModelElementsRedisRepository
	*PropertiesRedisRepository
	*UnitsRedisRepository
	*UsersRedisRepository
}

const (
	UnknownRole = "unknown" // TODO: remove?
	userPrefix  = "users"
)

func NewRedisStorage(host, port, password string) *RedisStorage {
	return &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		LayersRedisRepository:        NewLayersRedisRepository(host, port, password, "lr"),
		ModelsRedisRepository:        NewModelsRedisRepository(host, port, password, "mr"),
		ModelElementsRedisRepository: NewModelElementsRedisRepository(host, port, password, "mer"),
		PropertiesRedisRepository:    NewPropertiesRedisRepository(host, port, password, "pr"),
		UnitsRedisRepository:         NewUnitsRedisRepository(host, port, password, "unr"),
		UsersRedisRepository:         NewUsersRedisRepository(host, port, password, "usr"),
	}
}

func (rs *RedisStorage) OpenConn(request *entities.AuthRequest, _ context.Context) (storage.ConnDB, string, error) {
	key := compositeKey(userPrefix, request.Username)
	role, err := rs.client.Get(context.Background(), key).Result()
	switch {
	case err == redis.Nil:
		err = rs.AddUser(request.Username, interfaces.UserDTO{
			Name:     request.Username,
			Password: request.Password,
			Role:     UnknownRole,
		})
		if err != nil {
			return nil, "", err
		}
		return request.Username, UnknownRole, nil
	case err != nil:
		return nil, "", err
	default:
		return request.Username, role, nil
	}
}

func (rs *RedisStorage) CloseConn(conn storage.ConnDB) error {
	username, ok := conn.(string)
	if !ok {
		return errors.New("cannot get string from argument")
	}

	key := compositeKey(userPrefix, username)
	return rs.client.Del(context.Background(), key).Err()
}
