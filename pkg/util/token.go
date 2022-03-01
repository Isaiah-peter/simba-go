package utils

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Authuser() (c fiber.Handler) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusBadRequest).JSON("unauthorized")
		},
		SigningKey: []byte("my_secret_key"),
	})
}

func UseToken(c *fiber.Ctx) jwt.MapClaims {
	users := c.Locals("user").(*jwt.Token)
	token := users.Claims.(jwt.MapClaims)

	return token
}
