package controllers

import (
	"simba-clone/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "error while parsing data"})
	}
	tk := transaction.CreateTrancaction()
	transferTo, td := models.GetUserByEmail(transaction.EmailTo)
	transferfrom, fd := models.GetUserByEmail(transaction.EmailFrom)
	if transaction.CurrencyType == "USD" && transferfrom.DollarAcount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount

	}
	if transaction.CurrencyType == "EUR" && transferfrom.EuroAccount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount
	}
	if transaction.CurrencyType == "PUD" && transferfrom.PoundsAcount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount
	}

	td.Save(transferTo)
	fd.Save(transferfrom)
	return c.Status(fiber.StatusBadRequest).JSON(tk)
}

func GetAllTransaction(c *fiber.Ctx) error {
	tk := new(models.Transaction)
	db.Find(tk)
	return c.Status(fiber.StatusBadRequest).JSON(tk)
}
