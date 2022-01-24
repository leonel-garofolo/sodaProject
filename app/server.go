package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/enviroment"
)

type Server struct {
}

func (a *Server) Start() *enviroment.Context {
	database := enviroment.CreateConnection("localhost", 3060, "soda")
	context := &enviroment.Context{
		App:      a.setupFiber(),
		Database: database,
		Log:      enviroment.CreateLog(),
	}
	context.App.Static("/", "./public")
	Router(context)
	context.App.Listen(":3000")

	return context
}

func (a *Server) setupFiber() *fiber.App {
	return fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})
}
