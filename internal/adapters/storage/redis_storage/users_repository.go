package redis_storage

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/redis/go-redis/v9"
)

type UsersRedisRepository struct {
	client  *redis.Client
	keyType string
}

func NewUsersRedisRepository(host, port, password, keyType string) *UsersRedisRepository {
	return &UsersRedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		}),
		keyType: keyType,
	}
}

func (r *UsersRedisRepository) AddUser(_ storage.ConnDB, user interfaces.UserDTO) error {
	key := compositeKey(r.keyType, user.Name)
	return r.client.Set(context.Background(), key, user.Role, 0).Err()
}
