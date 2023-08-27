package main

import (
	"go-database/file"
	"go-database/routes"
	"go-database/scheduled"
)

func before() {

	createDirData()

	go routes.CreateRoutes()
	go scheduled.StartScheduled()
}

func createDirData() {
	if !file.DirExists("data/") {
		file.CreateDir("data/")
	}
}
