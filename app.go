package main

import (
	"github.com/leonel-garofolo/soda/cmd"
	"golang.org/x/mobile/app"
)

func main() {
	api := app.Server{}
	api.Start()
	//TODO when I have to execution?
	cmd.Execute()
}
