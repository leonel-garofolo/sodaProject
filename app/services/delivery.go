package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/repositories"
)

type DeliveryService struct {
	Dao *repositories.Dao
}

func (d *DeliveryService) GetClientForDelivery(c *fiber.Ctx) error {
	log.Println("entry-> GetClientForDelivery")
	return c.JSON(d.Dao.GetClientForDelivery())
}

func (d *DeliveryService) GetDeliveries(c *fiber.Ctx) error {
	log.Println("entry-> GetDeliveries")
	return c.JSON(d.Dao.GetDeliveries())
}
