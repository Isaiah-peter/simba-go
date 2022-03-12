package controllers

import (
	"simba-clone/pkg/models"
	utils "simba-clone/pkg/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing json"})
	}
	u := newUser.CreateUser()
	return c.Status(fiber.StatusOK).JSON(&u)
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing json"})
	}
	u := models.FindOne(user.Email, user.Password)
	return c.Status(fiber.StatusOK).JSON(&u)
}

func GetAllUser(c *fiber.Ctx) error {
	u := models.GetUser()
	return c.Status(fiber.StatusOK).JSON(&u)
}

func GetSingleUser(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "fail to convert to int"})
	}
	u, _ := models.GetUserById(id)
	return c.Status(fiber.StatusOK).JSON(u)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error  while parsing json"})
	}
	param := c.Params("id")
	id, _ := strconv.Atoi(param)
	u, db := models.GetUserById(id)
	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.Password != "" {
		hash := utils.HashPassword(user.Password)
		u.Password = hash
	}

	if user.DollarAcount != 0 {
		u.DollarAcount = user.DollarAcount
	}
	if user.EuroAccount != 0 {
		u.EuroAccount = user.EuroAccount
	}
	if user.PoundsAcount != 0 {
		u.PoundsAcount = user.PoundsAcount
	}
	db.Save(u)
	return c.Status(fiber.StatusOK).JSON(&u)
}
