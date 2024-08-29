package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"redis-go/database"
	"redis-go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app:= fiber.New()
	
    
	// Default CORS config allows all origins
    app.Use(cors.New())
	app.Use(logger.New())

	// Create a new Redis client
	rdb := database.ConnectToRedis()
	if rdb == nil {
		return
	}

	app.Get("/", func(c *fiber.Ctx) error {
		ctx := context.Background()
		cacheKey := "todos"

		// Try to get data from Redis
		cachedData, err := rdb.Get(ctx, cacheKey).Result()
		if err == nil {
			// Data found in Redis, return it
			var todos []models.Todo
			err = json.Unmarshal([]byte(cachedData), &todos)
			if err == nil {
				return c.JSON(todos)
			}
			// If unmarshaling fails, we'll fetch from API
		}
		// Make a GET request to the JSONPlaceholder API
		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch todos",
			})
		}
		fmt.Println("Response:", resp.Status)
		fmt.Println("Response:", resp.StatusCode)
		fmt.Println("Response:", resp.Body)
		
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read response body",
			})
		}

		// Parse the JSON response into a slice of Todo structs
		var todos []models.Todo
		err = json.Unmarshal(body, &todos)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse JSON",
			})
		}
		// Store in Redis
		jsonData, _ := json.Marshal(todos)
		err = rdb.Set(ctx, cacheKey, jsonData, 30*time.Second).Err() // Cache for 1 hour
		if err != nil {
			fmt.Println("Failed to cache data in Redis:", err)
		}

		// Return the todos as JSON
		return c.JSON(todos)
	})

		// Start the Fiber app
	fmt.Println("Server is running on :3000")
	app.Listen(":3000")

	// Close the Redis client
	// defer rdb.Close()

	//! String Operations
	// Get a key-value pair
	// name := operations.GetStringValue(rdb,"msg:1")
	// if name == "" {
	// 	fmt.Println("Error getting value")
	// }
	// fmt.Println("Name:", name)

	// Set a key-value pair
	// err := operations.SetStringValue(rdb,"msg:4", "Hello, World!")
	// if err != nil {
	// 	fmt.Println("Error setting value")
	// }

	// Expire a key
	// Experr := operations.ExpireKey(rdb,"msg:4")
	// if Experr != nil {
	// 	fmt.Println("Error setting value")
	// }
	// fmt.Println("Key will be expire in 10 secs")

	//! List Operations
	// lpush
	// err := operations.PushLeftToList(rdb, "list", "Hello")
	// err = operations.PushLeftToList(rdb, "list", "Hey")
	// err = operations.PushLeftToList(rdb, "list", "Good")
	// err = operations.PushLeftToList(rdb, "list", "Great")
	// if err != nil {
	// 	fmt.Println("Error pushing to list")
	// }

	// lpop
	// val := operations.PopLeftFromList(rdb, "list")
	// fmt.Println("Popped from list:", val)

	// blpop
	// val := operations.PopLeftWithBloackingStateFromList(rdb, "list",10)
	// fmt.Println("Popped from list:", val)
	
}