package main

import (
	"log"
	"os"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/eze-insurance/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	db := database.DbConnect()

	db.AutoMigrate(&model.AllRisk{}, &model.Marine{}, &model.GoodInTransit{}, &model.Shuttlers{})

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "System functioning well!"})
	})

	routes.Routes(app)

	log.Fatal(app.Listen(":" + port))
}
