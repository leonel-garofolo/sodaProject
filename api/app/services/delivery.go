package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/repositories"
	"github.com/leonel-garofolo/soda/app/utilities"
)

type DeliveryService struct {
	Dao *repositories.Dao
}

func (d *DeliveryService) GetClientForDelivery(c *fiber.Ctx) error {
	codRoot := utilities.ParseIntNoError(c.Query("cod", "-1"))
	log.Println("entry-> GetClientForDelivery-> ", codRoot)
	return c.JSON(d.Dao.GetClientForDelivery(codRoot))
}

func (d *DeliveryService) GetDeliveries(c *fiber.Ctx) error {
	log.Println("entry-> GetDeliveries")
	return c.JSON(d.Dao.GetDeliveries())
}
