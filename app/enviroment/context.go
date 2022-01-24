package enviroment

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
)

type Context struct {
	App      *fiber.App
	Database Database
	Log      string
}

type Database struct {
	Ip         string
	Port       int32
	Schema     string
	Connection *sql.DB
}

func CreateConnection(ip string, port int32, schemaName string) Database {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/soda")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Stablishment connection Mysql [ip:%s|port:%d|schemaName:%s]\n", ip, port, schemaName)
	return Database{
		Ip:         ip,
		Port:       port,
		Schema:     schemaName,
		Connection: db,
	}
}

func CreateLog() string {
	return ""
}
