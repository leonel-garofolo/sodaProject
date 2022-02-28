package reports

import (
	"testing"

	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/repositories"
)

func TestSheetRoot(t *testing.T) {
	daos := repositories.New(repositories.Dao{
		Database: enviroment.Database{
			Ip:     "localhost",
			Port:   3060,
			Schema: "soda",
		},
	})

	r := New(&Reports{
		Dao: daos,
	})

	r.BuildSheetRoot("10")
}
