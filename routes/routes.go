package routes

import (
	"github.com/eze-insurance/service"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	route := app.Group("/api/v1/")

	route.Post("/marine", service.CreateMarinePolicy)
	route.Post("/all-risk", service.CreateAllRiskPolicy)
	route.Post("/goods", service.CreateGoodsInTransitPolicy)
	route.Post("/shuttlers", service.CreateShuttlersPolicy)
	route.Post("/jokes", service.Jokes)
}
