package operations

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func GetStringValue(rdb *redis.Client, key string) string {
	// Create a context
	ctx := context.Background()
	// Get the value for the key
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return ""
	}

	return val
}

func SetStringValue(rdb *redis.Client, key string, value string) error {
	// Create a context
	ctx := context.Background()
	// Set a key-value pair
	err := rdb.Set(ctx, key,value, 0).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return err
	}
	return nil
}


func ExpireKey(rdb *redis.Client, key string) error {
	// Create a context
	ctx := context.Background()
	// Set a key-value pair
	err := rdb.Expire(ctx, key , 10*time.Second).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return err
	}
	return nil
}