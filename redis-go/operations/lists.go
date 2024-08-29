package operations

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func PushLeftToList(rdb *redis.Client, key string, value string) error {
	// Create a context
	ctx := context.Background()
	// Push a value to the list
	err := rdb.LPush(ctx, key, value).Err()
	if err != nil {
		fmt.Println("Error pushing to list:", err)
		return err
	}
	return nil
}

func PushRightToList(rdb *redis.Client, key string, value string) error {
	// Create a context
	ctx := context.Background()
	// Push a value to the list
	err := rdb.LPush(ctx, key, value).Err()
	if err != nil {
		fmt.Println("Error pushing to list:", err)
		return err
	}
	return nil
}


func PopLeftFromList(rdb *redis.Client, key string) string {
	// Create a context
	ctx := context.Background()
	// Pop a value from the list
	val, err := rdb.LPop(ctx, key).Result()
	if err != nil {
		fmt.Println("Error popping from list:", err)
		return ""
	}
	return val
}


func PopRightFromList(rdb *redis.Client, key string) string {
	// Create a context
	ctx := context.Background()
	// Pop a value from the list
	val, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		fmt.Println("Error popping from list:", err)
		return ""
	}
	return val
}

func PopLeftWithBloackingStateFromList(rdb *redis.Client, key string, timeInSeconds int) string {
	// Create a context
	ctx := context.Background()
	// Convert time to duration in seconds
	duration := time.Duration(timeInSeconds) * time.Second
	// Pop a value from the list
	val, err := rdb.BLPop(ctx, duration, key).Result()
	if err != nil {
		fmt.Println("Error popping from list:", err)
		return ""
	}
	fmt.Println("Popped from list:", val)
	return val[1] // Assuming you want to return the second element of the result slice
}