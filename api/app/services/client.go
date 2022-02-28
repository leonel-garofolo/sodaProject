package services

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/model"
	"github.com/leonel-garofolo/soda/app/repositories"
)

type ClientService struct {
	Dao *repositories.Dao
}

func (s *ClientService) Save(c *fiber.Ctx) error {
	client := new(model.Client)
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := model.ValidateStruct(*client)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.JSON(s.Dao.SaveClient(client))
}

func (s *ClientService) Delete(c *fiber.Ctx) error {
	idClient := c.Query("id", "-1")
	idClientInt, err := strconv.Atoi(idClient)
	if err != nil {
		return c.JSON(nil)
	}
	return c.JSON(s.Dao.DeleteClient(idClientInt))
}
