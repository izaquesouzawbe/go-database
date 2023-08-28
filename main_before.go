package main

import (
	"go-database/commands"
	"go-database/reading"
	"go-database/routes"
	"go-database/scheduled"
)

func before() {

	commands.CreateDirData()

	go routes.CreateRoutes()
	go scheduled.StartScheduled()

	reading.ReadInput()

}
