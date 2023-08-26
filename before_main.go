package main

import (
	"go-database/file"
	"go-database/routes"
	"go-database/scheduleds"
)

func before() {

	createDirData()

	routes.CreateRoutes()
	go scheduleds.ScheduledTask()
}

func createDirData() {
	if !file.DirExists("data/") {
		file.CreateDir("data/")
	}
}
