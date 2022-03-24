package service

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/gofiber/fiber/v2"
)

func CreateMarinePolicy(c *fiber.Ctx) error {

	marine := new(model.MarineRequest)

	err := c.BodyParser(marine)

	if err != nil {
		log.Fatal(err)
	}

	var invoice float64

	if marine.Currency == "ngn" {
		amount := marine.Amount / 100
		invoice = float64(amount) + (0.1 * float64(marine.Amount))
	} else {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  false,
			"message": "Currency is not naira"})
	}

	ratePercentage := 0.2

	rate := ratePercentage / 100 * invoice

	marineStruct := model.Marine{
		Value:          marine.Amount,
		TotalInvoice:   invoice,
		Rate:           float64(rate),
		RatePercentage: ratePercentage,
	}

	db := database.DbConnect()

	result := db.Create(&marineStruct)

	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Created New Marine Policy",
		"policy":  marineStruct,
		"status":  true})
}
