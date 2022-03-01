package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/services"
)

const app_PATH = "/app"

type Router struct {
	App             *fiber.App
	DeliveryService services.DeliveryService
	ClientService   services.ClientService
	ReportService   services.ReportService
}

func New(router Router) *Router {
	return &router
}

func (r *Router) Setup() {
	log.Println("setup -> Router config")
	app := r.App.Group(app_PATH)
	app.Get("/delivery", r.DeliveryService.GetDeliveries)
	app.Get("/deliveryCode", r.DeliveryService.GetDeliveriesCode)
	app.Get("/deliveriesToClient", r.DeliveryService.GetClientForDelivery)
	app.Post("/client", r.ClientService.Save)
	app.Delete("/client", r.ClientService.Delete)
	app.Get("/report", r.ReportService.GenerateReport)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(r.App.Stack())
	})
}
