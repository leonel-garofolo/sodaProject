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
	log.Println("Save")
	client := &model.Client{
		Id:           1,
		Address:      "",
		NumAddress:   123,
		Order:        1,
		PricePerSoda: 1,
		PricePerBox:  2,
		IdDelivery:   1,
		IdRoot:       1,
		Due:          1,
	}
	if client.Id > 0 {
		s.update(client)
	} else {
		client.Id = s.insert(client)
	}
	return c.JSON(client)
}

func (s *ClientService) Delete(idClient int) {
	log.Println("Delete")
	sqlStatement := `
	DELETE FROM client where id_client $1
	`
	_, err := s.Context.Database.Connection.Exec(sqlStatement, idClient)
	if err != nil {
		panic(err)
	}
}

func (s *ClientService) insert(client *model.Client) int {
	sqlStatement := `
	INSERT INTO client(address, number, order, id_delivery, id_root, price_per_soda, price_per_box)
	VALUE($1, $2, $3, $4, $5, $6, $7)
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
	sqlStatement := `
	UPDATE client
	SET address= $2, number= $3, order= $4, id_delivery= $5, id_root= $6, 
		price_per_soda=$7, price_per_box= $8
	WHERE id_client = $1`
	_, err := s.Context.Database.Connection.Exec(sqlStatement,
		client.Id,
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
}
