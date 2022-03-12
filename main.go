package main

import (
	"os"

	"simba-clone/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE,HEAD,PUT",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	routes.UserRoute(app)
	routes.Transaction(app)
	app.Listen(":" + port)
}
