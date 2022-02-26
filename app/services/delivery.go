package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/enviroment"
	model "github.com/leonel-garofolo/soda/app/model"
)

type DeliveryService struct {
	Context *enviroment.Context
}

func (d *DeliveryService) GetClientForDelivery(c *fiber.Ctx) error {
	log.Println("entry-> GetClientForDelivery")
	db := d.Context.Database.Connection
	rows, err := db.Query("select * from client")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var clients []model.Client
	for rows.Next() {
		var client model.Client
		if err := rows.Scan(&client.Id, &client.Address, &client.NumAddress, &client.Order, &client.IdDelivery, &client.IdRoot, &client.PricePerSoda, &client.PricePerBox); err != nil {
			return c.JSON(nil)
		}
		clients = append(clients, client)
	}
	if err = rows.Err(); err != nil {
		return c.JSON(nil)
	}
	return c.JSON(clients)
}

func (d *DeliveryService) GetDeliveries(c *fiber.Ctx) error {
	log.Println("entry-> GetDeliveries")

	db := d.Context.Database.Connection
	rows, err := db.Query("select * from delivery")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var deliveries []model.Delivery
	for rows.Next() {
		var delivery model.Delivery
		if err := rows.Scan(&delivery.Id, &delivery.Name); err != nil {
			return c.JSON(nil)
		}

		deliveryRoot := getDeliveryRoot(db, delivery.Id)
		for i := 0; i < len(deliveryRoot); i++ {
			delivery.Code = append(delivery.Code, deliveryRoot[i].Code)
		}
		deliveries = append(deliveries, delivery)
	}
	return c.JSON(deliveries)
}

func getDeliveryRoot(db *sql.DB, idDelivery int) []model.DeliveryRoot {
	log.Println("entry-> getDeliveryRoot")
	rows, err := db.Query("select * from delivery_root dr where dr.id_delivery= ?", idDelivery)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var deliveriesRoot []model.DeliveryRoot
	for rows.Next() {
		var deliveryRoot model.DeliveryRoot
		if err := rows.Scan(&deliveryRoot.IdDelivery, &deliveryRoot.IdRoot, &deliveryRoot.Code); err != nil {
			return deliveriesRoot
		}
		deliveriesRoot = append(deliveriesRoot, deliveryRoot)
	}
	fmt.Printf("number of Root finded to delivery: %d", len(deliveriesRoot))
	fmt.Println()
	return deliveriesRoot
}
