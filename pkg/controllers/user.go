package controllers

import (
	"simba-clone/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

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
	u := models.FindUser(user.Email, user.Password)
	return c.Status(fiber.StatusOK).JSON(&u)
}

func GetAllUser(c *fiber.Ctx) error {
	allUser := new(models.User)
	db.Find(&allUser)
	return c.Status(fiber.StatusOK).JSON(allUser)
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
	db.Save(u)
	return c.Status(fiber.StatusOK).JSON(&u)
}
