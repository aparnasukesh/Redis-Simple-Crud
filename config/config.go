package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
	dsn := "host=localhost user=postgres password=2585 dbname=redis_sample port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Connected to PostgresSQL")
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
		DB:   0,                // Redis Database
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to conntect to Redis : %v ", err)
		os.Exit(1)
	}
	fmt.Println("Connected to Redis")
}
