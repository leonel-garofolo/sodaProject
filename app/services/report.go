package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/reports"
)

type ReportService struct {
	Reports *reports.Reports
}

func (r *ReportService) GenerateReport(c *fiber.Ctx) error {
	idRoot := c.Query("id", "-1")
	log.Println("generateReport root-> ", idRoot)
	r.Reports.BuildSheetRoot(idRoot)
	return c.JSON(idRoot)
}
