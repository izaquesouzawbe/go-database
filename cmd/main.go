package main

import (
	"go-zdb-api/internal/commands"
	"go-zdb-api/internal/routes"
	"go-zdb-api/internal/scheduled"
)

func main() {

	before()

}

func before() {

	commands.CreateDirData()

	go routes.CreateRoutes()
	go scheduled.StartScheduled()

	select {}
	//reading.ReadInput()

}
