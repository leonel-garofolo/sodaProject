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
	"fmt"
	"log"

	"github.com/Kirides/go-dbf"
	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/model"
	"github.com/leonel-garofolo/soda/app/repositories"
	"github.com/leonel-garofolo/soda/app/utilities"
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
		cmd.Printf("%s %s %s %s", Prefix, Root, Migrate, "<PATH>/<file_name>.DBF")
		cmd.Println("")

		//TODO validation file path

		//TODO call method

	},
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

func ProcessMigrateClientData(filePath string, codRoot int) {
	dbfDB, err := dbf.Open(filePath, charmap.Windows1252.NewDecoder())
	if err != nil {
		panic(err)
	}
	defer dbfDB.Close()

	daos := repositories.New(repositories.Dao{
		Database: enviroment.Database{
			Ip:     "localhost",
			Port:   3060,
			Schema: "soda",
		},
	})
	idDelivery, idRoot, errRoot := daos.GetIdRoot(codRoot)
	if errRoot != nil {
		log.Println("CodRoot not found")
		return
	}
	log.Println("count rows: ", dbfDB.CalculatedRecordCount())
	var parseOption dbf.ParseOption
	proccess := func(r *dbf.Record) error {
		if !r.Deleted() {
			data, err := r.ToMap() // returns a map[string]interface{}
			if err != nil {
				panic(err)
			}
			//Example: map[CALLE:Liniers         CODIGO:1 DEUDA:280 NUMERO:1463 NUMEROREPA:0 PRECIO:4 PRECIO2:0]
			log.Println(data)
			daos.SaveClient(&model.Client{
				Order:        utilities.ParseIntNoError(fmt.Sprintf("%d", data["CODIGO"])),
				Address:      fmt.Sprintf("%s", data["CALLE"]),
				NumAddress:   utilities.ParseIntNoError(fmt.Sprintf("%d", data["NUMERO"])),
				Debt:         utilities.ParseFloatNoError(fmt.Sprintf("%.2f", data["DEUDA"])),
				PricePerSoda: utilities.ParseFloatNoError(fmt.Sprintf("%.2f", data["PRECIO"])),
				PricePerBox:  utilities.ParseFloatNoError(fmt.Sprintf("%.2f", data["PRECIO2"])),
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
