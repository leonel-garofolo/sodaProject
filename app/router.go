package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/services"
)

const app_PATH = "/app"

func Router(context *enviroment.Context) {
	log.Println("setup -> Router config")
	app := context.App.Group(app_PATH)
	deliveryService := services.DeliveryService{
		Context: context,
	}
	clientService := services.ClientService{
		Context: context,
	}
	reportService := services.Report{
		Context: context,
	}

	app.Get("/delivery", deliveryService.GetDeliveries)
	app.Get("/deliveriesToClient", deliveryService.GetClientForDelivery)
	app.Post("/client", clientService.Save)
	app.Delete("/client", clientService.Delete)
	app.Get("/report", reportService.GenerateReport)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(context.App.Stack())
	})
}
