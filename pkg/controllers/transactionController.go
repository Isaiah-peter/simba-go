package controllers

import (
	"simba-clone/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "error while parsing data"})
	}
	transferTo, td := models.GetUserByEmail(transaction.EmailTo)
	transferfrom, fd := models.GetUserByEmail(transaction.EmailFrom)
	if transaction.CurrencyType == "USD" && transferfrom.DollarAcount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount
		transaction.Status = "Succsessful"
	}
	if transaction.CurrencyType == "EUR" && transferfrom.EuroAccount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount
		transaction.Status = "Succsessful"
	}
	if transaction.CurrencyType == "GBP" && transferfrom.PoundsAcount <= transaction.Amount {
		transferTo.DollarAcount = transferTo.DollarAcount + transaction.Amount
		transferfrom.DollarAcount = transferfrom.DollarAcount - transaction.Amount
		transaction.Status = "Succsessful"
	}
	if transferfrom.DollarAcount > transaction.Amount {
		transaction.Status = "Pending"
	}
	if transferfrom.EuroAccount > transaction.Amount {
		transaction.Status = "Pending"
	}
	if transferfrom.PoundsAcount > transaction.Amount {
		transaction.Status = "Pending"
	}
	tk := transaction.CreateTrancaction()
	td.Save(transferTo)
	fd.Save(transferfrom)
	return c.Status(fiber.StatusBadRequest).JSON(tk)
}

func GetAllTransaction(c *fiber.Ctx) error {
	tk := new(models.Transaction)
	db.Find(tk)
	return c.Status(fiber.StatusBadRequest).JSON(tk)
}

func UpdateTransaction(c *fiber.Ctx) error {
	tk := new(models.Transaction)
	if err := c.BodyParser(tk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "fail to parse"})
	}
	transctionId := c.Params("id")
	id, err := strconv.Atoi(transctionId)
	if err != nil {
		panic(err)
	}
	transaction, db := models.GetTransactionById(id)
	db.Save(transaction)
	return c.Status(fiber.StatusBadRequest).JSON(transaction)
}
