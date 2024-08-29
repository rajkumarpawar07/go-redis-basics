package database

import (
    "context"
    "fmt"
    redis "github.com/redis/go-redis/v9"
)


func ConnectToRedis() *redis.Client {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Create a context
	ctx := context.Background()

	// Ping the Redis server
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}

	fmt.Println("Connected to Redis:", pong)
	return rdb
}
  