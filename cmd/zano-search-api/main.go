package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getEnvVar(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

func main() {
	//mongoUri := getEnvVar("MONGO_URI_DEV")

	// Initialize a new Fiber app
	// Custom config
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Bazaar",
		AppName:       "Bazaar API v0.1.0",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	// Middleware

	// Initialize custom limiter config using MongoDB connection string
	//storage := mongodb.New(mongodb.Config{
	//	ConnectionURI: mongoUri,
	//	Database:      "fiber",
	//	Collection:    "fiber_storage",
	//	Reset:         false,
	//})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://zanobazaar.com, https://bazaargo.com",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(recover.New())
	app.Use(helmet.New())
	//app.Use(limiter.New(limiter.Config{
	//	Storage: storage,
	//}))

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})
	// Get alias
	v1.Post("/search", func(c fiber.Ctx) error {
		//o := new(Offers)
		//if err := c.Bind().JSON(o); err != nil {
		//	return err
		//}

		// Send a string response to the client
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
		})
	})

	// Start the server on port 3000
	// Custom config
	_ = app.Listen(":8080", fiber.ListenConfig{
		EnablePrefork:         false,
		DisableStartupMessage: false,
	})
}
