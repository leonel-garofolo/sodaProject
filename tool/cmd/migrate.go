/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	// with go modules enabled (GO111MODULE=on or outside GOPATH)
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Kirides/go-dbf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/charmap"
)

const (
	Migrate = "migrate"
)

// migrate represents the get command
var migrate = &cobra.Command{
	Use:   Migrate,
	Short: "Migrate data of xBase to Mysql database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Printf("Error in commmand execution, please use the follow formatter /n %s %s %s %s", Prefix, Root, Migrate, "<PATH>/<file_name>.DBF")
			cmd.Println("")
		} else {
			path := args[0]
			code := args[1]
			codeNum, err := strconv.Atoi(code)
			if err == nil {
				if isFilePathValid(path) {
					ProcessMigrateClientData(path, codeNum)
				}
			} else {
				cmd.Println("Oops! Error parser code root")
			}
		}
	},
}

func isFilePathValid(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

type ClientDBF struct {
	Codigo    string
	Precio    string
	Direccion string
	Numero    string
	Deuda     string
	Precio2   string
	Reparto   string
}

type Client struct {
	Id           int     `json:"id"`
	Order        int     `json:"order" validate:"required,min=1,max=999"`
	Address      string  `json:"address" validate:"required,min=3,max=100"`
	NumAddress   int     `json:"num_address"  validate:"required,min=1,max=99999"`
	PricePerSoda float64 `json:"price_per_soda"  validate:"required,min=0.1,max=99999.99"`
	PricePerBox  float64 `json:"price_per_box" validate:"min=0.1,max=99999.99"`
	Debt         float64 `json:"debt" validate:"min=0.1,max=99999999.99"`
	IdDelivery   int     `json:"id_delivery" validate:"required"`
	IdRoot       int     `json:"id_root" validate:"required"`
}

func ParseIntNoError(s string) int {
	f, _ := strconv.Atoi(s)
	return f
}
func ParseFloatNoError(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func ProcessMigrateClientData(filePath string, codRoot int) {
	dbfDB, err := dbf.Open(filePath, charmap.Windows1252.NewDecoder())
	if err != nil {
		panic(err)
	}
	defer dbfDB.Close()

	db := createConnection()
	log.Println("count rows: ", dbfDB.CalculatedRecordCount())

	idDelivery, idRoot, errorCode := getDeliveryAndRoot(db, codRoot)
	if errorCode != nil {
		panic(errorCode)
	}

	var parseOption dbf.ParseOption
	proccess := func(r *dbf.Record) error {
		if !r.Deleted() {
			data, err := r.ToMap() // returns a map[string]interface{}
			if err != nil {
				panic(err)
			}
			//Example: map[CALLE:Liniers         CODIGO:1 DEUDA:280 NUMERO:1463 NUMEROREPA:0 PRECIO:4 PRECIO2:0]
			log.Println(data)
			insert(db, &Client{
				Order:        ParseIntNoError(fmt.Sprintf("%d", data["CODIGO"])),
				Address:      fmt.Sprintf("%s", data["CALLE"]),
				NumAddress:   ParseIntNoError(fmt.Sprintf("%d", data["NUMERO"])),
				Debt:         ParseFloatNoError(fmt.Sprintf("%.2f", data["DEUDA"])),
				PricePerSoda: ParseFloatNoError(fmt.Sprintf("%.2f", data["PRECIO"])),
				PricePerBox:  ParseFloatNoError(fmt.Sprintf("%.2f", data["PRECIO2"])),
				IdDelivery:   idDelivery,
				IdRoot:       idRoot,
			})
		}
		return nil
	}
	err = dbfDB.Scan(proccess, parseOption)

	if err != nil {
		panic(err)
	}
}

func GetMigration() *cobra.Command {
	return migrate
}

func init() {
	rootCmd.AddCommand(GetMigration())
	migrate.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createConnection() *sql.DB {
	sConnection := fmt.Sprintf("root:1234@tcp(%s:%d)/%s", "localhost", 3306, "soda")
	db, err := sql.Open("mysql", sConnection)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Stablishment connection Mysql -> %s\n", sConnection)
	return db
}

func insert(db *sql.DB, client *Client) {
	log.Println("insert -> ", &client)

	sqlStatement := `
	INSERT INTO client(address, address_number, num_order, id_delivery, id_root, price_per_soda, price_per_box, debt)
	VALUE(?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(sqlStatement,
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
}

func getDeliveryAndRoot(db *sql.DB, codeRoot int) (int, int, error) {
	log.Println("getDeliveryAndRoot-> ", codeRoot)

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
