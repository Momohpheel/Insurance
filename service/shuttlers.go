package service

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/gofiber/fiber/v2"
)

func CreateShuttlersPolicy(c *fiber.Ctx) error {
	shuttlers := new(model.ShuttlersRequest)

	err := c.BodyParser(shuttlers)

	if err != nil {
		log.Fatal(err)
	}

	shuttlerStruct := model.Shuttlers{
		FullName:         shuttlers.FullName,
		Address:          shuttlers.Address,
		Business:         shuttlers.Business,
		MaritalStatus:    shuttlers.MaritalStatus,
		DateOfBirth:      shuttlers.DateOfBirth,
		PhoneNumber:      shuttlers.PhoneNumber,
		Email:            shuttlers.Email,
		NextOfKin:        shuttlers.NextOfKin,
		NextOfKinAddress: shuttlers.NextOfKin,
		NextOfKinPhone:   shuttlers.NextOfKinPhone,
	}

	db := database.DbConnect()

	result := db.Create(&shuttlerStruct)

	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Created New Shuttlers Policy",
		"policy":  shuttlerStruct,
		"status":  true})
}
