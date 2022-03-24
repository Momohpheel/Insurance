package service

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/gofiber/fiber/v2"
)

func CreateGoodsInTransitPolicy(c *fiber.Ctx) error {
	goods := new(model.GoodsInTransitRequest)

	err := c.BodyParser(goods)

	if err != nil {
		log.Fatal(err)
	}

	ratePercentage := 0.5

	rate := ratePercentage / 100 * float64(goods.EstimatedAnnualCarrying)

	goodsStruct := model.GoodInTransit{
		Name:                    goods.Name,
		Address:                 goods.Address,
		Occupation:              goods.Occupation,
		Phone:                   goods.Phone,
		Email:                   goods.Email,
		TypeOfGoods:             goods.TypeOfGoods,
		NatureOfGoods:           goods.NatureOfGoods,
		LimitAnyoneCarrying:     goods.LimitAnyoneCarrying,
		EstimatedAnnualCarrying: goods.EstimatedAnnualCarrying,
		ModeOfTransaporting:     goods.ModeOfTransaporting,
		OwnVehicles:             goods.OwnVehicles,
		RegisterationNumber:     goods.RegisterationNumber,
		Rate:                    rate,
		RatePercentage:          ratePercentage,
	}

	db := database.DbConnect()

	result := db.Create(&goodsStruct)

	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Created New Goods In Transit Policy",
		"policy":  goodsStruct,
		"status":  true})
}
