package enviroment

import (
	"database/sql"
	"log"

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

type Config struct {
	Database Database
}

func New(config Config) *Context {
	context := new(Context)
	context.Database = config.Database
	return context
}

func (c *Context) Setup() {
	log.Println("Context setup")
	c.setupFiber()
	c.createLog()
}

func (c *Context) setupFiber() {
	c.App = fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})
	c.App.Static("/", "./public")
}

func (c *Context) createLog() string {
	return ""
}
