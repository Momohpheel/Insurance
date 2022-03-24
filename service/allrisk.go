package service

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/gofiber/fiber/v2"
)

func CreateAllRiskPolicy(c *fiber.Ctx) error {
	allrisk := new(model.AllRiskRequest)

	err := c.BodyParser(allrisk)

	if err != nil {
		log.Fatal(err)
	}

	ratePercentage := 5.0

	rate := ratePercentage / 100 * float64(allrisk.SumInsured)

	allriskStruct := model.AllRisk{
		Name:           allrisk.Name,
		SerialNumber:   allrisk.SerialNumber,
		Configuration:  allrisk.Configuration,
		SumInsured:     allrisk.SumInsured,
		Rate:           rate,
		RatePercentage: ratePercentage,
	}

	db := database.DbConnect()

	result := db.Create(&allriskStruct)

	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Created New All Risk Policy",
		"policy":  allriskStruct,
		"status":  true})
}
