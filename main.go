package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	app := fiber.New()
	app.Use(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE,HEAD,PUT",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	log.Fatal(app.Listen(port))
}
