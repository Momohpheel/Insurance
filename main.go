package main

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/eze-insurance/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	db := database.DbConnect()

	db.AutoMigrate(&model.AllRisk{}, &model.Marine{}, &model.GoodInTransit{}, &model.Shuttlers{})

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "System functioning well!"})
	})

	routes.Routes(app)

	log.Fatal(app.Listen(":8000"))
}
