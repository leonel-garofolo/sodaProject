package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/model"
)

type Dao struct {
	Database enviroment.Database
}

func New(dao Dao) *Dao {
	dao.createConnection()
	return &dao
}

func (c *Dao) createConnection() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/soda")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	c.Database.Connection = db
	fmt.Printf("Stablishment connection Mysql [ip:%s|port:%d|schemaName:%s]\n", c.Database.Ip, c.Database.Port, c.Database.Schema)
}

func (d *Dao) GetClientForDelivery(codRoot int) []model.Client {
	db := d.Database.Connection
	sqlStatement := `
		select c.*  
		from client c 
		inner join delivery_root dr on dr.id_delivery = c.id_delivery and dr.id_root = c.id_root 
		where dr.code = ?`
	rows, err := db.Query(sqlStatement, codRoot)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var clients []model.Client
	if err = rows.Err(); err != nil {
		return clients
	}

	for rows.Next() {
		var client model.Client
		if err := rows.Scan(&client.Id, &client.Address, &client.NumAddress, &client.Order, &client.IdDelivery, &client.IdRoot, &client.PricePerSoda, &client.PricePerBox, &client.Debt); err != nil {
			break
		}
		clients = append(clients, client)
	}
	return clients
}

func (d *Dao) GetDeliveries() []model.Delivery {
	db := d.Database.Connection
	rows, err := db.Query("select * from delivery")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var deliveries []model.Delivery
	for rows.Next() {
		var delivery model.Delivery
		if err := rows.Scan(&delivery.Id, &delivery.Name); err != nil {
			break
		}

		deliveryRoot := d.getDeliveryRoot(delivery.Id)
		for i := 0; i < len(deliveryRoot); i++ {
			delivery.Code = append(delivery.Code, deliveryRoot[i].Code)
		}
		deliveries = append(deliveries, delivery)
	}

	return deliveries
}

func (d *Dao) getDeliveryRoot(idDelivery int) []model.DeliveryRoot {
	log.Println("entry-> getDeliveryRoot")
	rows, err := d.Database.Connection.Query("select * from delivery_root dr where dr.id_delivery= ?", idDelivery)
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

func (d *Dao) SaveClient(client *model.Client) *model.Client {
	if client.Id > 0 {
		d.update(client)
	} else {
		client.Id = d.insert(client)
	}
	return client
}

func (d *Dao) DeleteClient(idClient int) bool {
	log.Println("delete-> ", idClient)
	sqlStatement := `
	DELETE FROM client where id_client =?
	`
	_, err := d.Database.Connection.Exec(sqlStatement, idClient)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (d *Dao) insert(client *model.Client) int {
	log.Println("insert -> ", &client)

	sqlStatement := `
	INSERT INTO client(address, number, num_order, id_delivery, id_root, price_per_soda, price_per_box, debt)
	VALUE(?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := d.Database.Connection.Exec(sqlStatement,
		client.Address,
		client.NumAddress,
		client.Order,
		client.IdDelivery,
		client.IdRoot,
		client.PricePerSoda,
		client.PricePerBox,
		client.Debt,
	)
	if err != nil {
		panic(err)
	}
	// TODO devolver ID
	return 1
}

func (d *Dao) update(client *model.Client) {
	log.Println("update -> ", &client)
	sqlStatement := `
	UPDATE client
	SET address= ?, number= ?, num_order= ?, id_delivery= ?, id_root= ?, 
		price_per_soda=?, price_per_box= ?
	WHERE id_client = ?`
	_, err := d.Database.Connection.Exec(sqlStatement,
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

func (d *Dao) GetClientsRoot(idRoot string) ([]interface{}, error) {
	idRootInt, errorParser := strconv.Atoi(idRoot)
	if errorParser != nil {
		panic(errorParser.Error())
	}

	db := d.Database.Connection
	rows, err := db.Query("select c.num_order, c.address, c.`number`, c.price_per_soda, c.price_per_box, debt from client c where c.id_root= ? order by c.num_order asc", idRootInt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var clients []interface{}
	for rows.Next() {
		var client model.Client
		if err := rows.Scan(
			&client.Order,
			&client.Address,
			&client.NumAddress,
			&client.PricePerSoda,
			&client.PricePerBox,
			&client.Debt,
		); err != nil {
			return nil, err
		}

		clients = append(clients, []string{
			strconv.Itoa(client.Order),
			fmt.Sprintf("%.2f", client.PricePerSoda),
			fmt.Sprintf("%.2f", client.PricePerBox),
			client.Address,
			strconv.Itoa(client.NumAddress),
			fmt.Sprintf("%.2f", client.Debt),
		})
	}
	return clients, nil
}

func (d *Dao) GetNameCodeRoot(idRoot string) (string, string, error) {
	idRootInt, errorParser := strconv.Atoi(idRoot)
	if errorParser != nil {
		panic(errorParser.Error())
	}

	db := d.Database.Connection
	rows, err := db.Query("select d.name, dr.code from soda.delivery_root dr inner join soda.delivery d on d.id_delivery = dr.id_delivery where dr.id_root= ?", idRootInt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var name string
	var code string
	if rows.Next() {
		var codeInt int
		if err := rows.Scan(
			&name,
			&codeInt,
		); err != nil {
			return "", "", err
		}
		code = strconv.Itoa(codeInt)
	}
	return name, code, nil
}

func (d *Dao) GetIdRoot(codeRoot int) (int, int, error) {
	db := d.Database.Connection
	rows, err := db.Query("select dr.id_delivery, dr.id_root from delivery_root dr where dr.code = ?", codeRoot)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var idDelivery int
	var idRoot int
	if rows.Next() {
		if err := rows.Scan(
			&idDelivery,
			&idRoot,
		); err != nil {
			return -1, -1, err
		}
	}
	return idDelivery, idRoot, nil
}
