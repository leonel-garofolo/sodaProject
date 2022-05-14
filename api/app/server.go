package app

import (
	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/reports"
	"github.com/leonel-garofolo/soda/app/repositories"
	"github.com/leonel-garofolo/soda/app/services"
)

type App struct {
}

func (a *App) Start() {
	context := enviroment.New(enviroment.Config{
		Database: enviroment.Database{
			Ip:     "db",
			Port:   3306,
			Schema: "soda",
		},
	})
	context.Setup()
	daos := a.setupDao(context.Database)
	reports := a.setupReports(daos)
	a.setupService(context, daos, reports)
	context.App.Listen(":3000")
}

func (a *App) setupDao(database enviroment.Database) *repositories.Dao {
	return repositories.New(repositories.Dao{
		Database: database,
	})
}

func (a *App) setupReports(daos *repositories.Dao) *reports.Reports {
	return reports.New(&reports.Reports{
		Dao: daos,
	})
}

func (a *App) setupService(context *enviroment.Context, daos *repositories.Dao, report *reports.Reports) {
	router := New(Router{
		App: context.App,
		DeliveryService: services.DeliveryService{
			Dao: daos,
		},
		ClientService: services.ClientService{
			Dao: daos,
		},
		ReportService: services.ReportService{
			Reports: report,
		},
	})
	router.Setup()
}
