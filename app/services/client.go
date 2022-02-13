package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/model"
)

type ClientService struct {
	Context *enviroment.Context
}

func (s *ClientService) Save(c *fiber.Ctx) error {
	log.Println("save proccesing")
	client := &model.Client{}
	if err := c.BodyParser(&client); err != nil {
		log.Println(err)
	}
	if client.Id > 0 {
		s.update(client)
	} else {
		client.Id = s.insert(client)
	}
	return c.JSON(client)
}

func (s *ClientService) Delete(c *fiber.Ctx) error {
	idClient := c.Query("id", "-1")
	log.Println("delete-> ", idClient)
	sqlStatement := `
	DELETE FROM client where id_client =?
	`
	_, err := s.Context.Database.Connection.Exec(sqlStatement, idClient)
	if err != nil {
		panic(err)
	}
	return c.JSON(idClient)
}

func (s *ClientService) insert(client *model.Client) int {
	log.Println("insert -> ", &client)

	sqlStatement := `
	INSERT INTO client(address, number, num_order, id_delivery, id_root, price_per_soda, price_per_box)
	VALUE(?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.Context.Database.Connection.Exec(sqlStatement,
		client.Address,
		client.NumAddress,
		client.Order,
		client.IdDelivery,
		client.IdRoot,
		client.PricePerSoda,
		client.PricePerBox,
	)
	if err != nil {
		panic(err)
	}
	// TODO devolver ID
	return 1
}

func (s *ClientService) update(client *model.Client) {
	log.Println("update -> ", &client)
	sqlStatement := `
	UPDATE client
	SET address= ?, number= ?, num_order= ?, id_delivery= ?, id_root= ?, 
		price_per_soda=?, price_per_box= ?
	WHERE id_client = ?`
	_, err := s.Context.Database.Connection.Exec(sqlStatement,
		client.Address,
		client.NumAddress,
		client.Order,
		client.IdDelivery,
		client.IdRoot,
		client.PricePerSoda,
		client.PricePerBox,
		client.Id,
	)
	if err != nil {
		panic(err)
	}
}
