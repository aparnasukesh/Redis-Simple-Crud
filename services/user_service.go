package services

import (
	"context"
	"encoding/json"
	"fmt"
	"redis_sample_project/config"
	"redis_sample_project/models"
	"redis_sample_project/repository"
)

func CreateUser(user models.User) error {
	return repository.CreateUser(user)
}

func GetUserByID(id uint) (models.User, error) {
	// Check if the user exists in redis cache
	redisKey := fmt.Sprintf("user_%d", id)
	val, err := config.RedisClient.Get(context.Background(), redisKey).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(val), &user)
		return user, nil
	}

	// If not in Redis,fetch from PostgresSQL
	user, err := repository.GetUserByID(id)
	if err == nil {
		// Store the user in Redis
		userJson, _ := json.Marshal(user)
		config.RedisClient.Set(context.Background(), redisKey, userJson, 0)
	}
	return user, err
}

func UpdateUser(user models.User) error {
	redisKey := fmt.Sprintf("user_%d", user.ID)
	config.RedisClient.Del(context.Background(), redisKey) //Invalid Redis cache
	return repository.UpdateUser(user)
}

func DeleteUser(id uint) error {
	redisKey := fmt.Sprintf("user_%d", id)
	config.RedisClient.Del(context.Background(), redisKey)
	return repository.DeleteUser(id)
}

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}
