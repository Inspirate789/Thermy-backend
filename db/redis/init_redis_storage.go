package main

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/redis_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"log"
	"os"
)

func main() {
	redisStorage := redis_storage.NewRedisStorage(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
	)
	err := redisStorage.AddUser(nil, interfaces.UserDTO{
		Name:     os.Getenv("POSTGRES_ADMIN_USERNAME"),
		Password: os.Getenv("POSTGRES_ADMIN_PASSWORD"),
		Role:     "admin",
	})
	if err != nil {
		log.Fatal(err)
	}
}
