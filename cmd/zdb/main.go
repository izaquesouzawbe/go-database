package main

import (
	"fmt"
	"go-zdb-api/internal/commands"
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/routes"
	"go-zdb-api/internal/scheduled"
)

func main() {

	before()

}

func before() {

	commands_func.CreateDirData()
	createDataTest()

	global.LoadDatabaseInMemory()

	go routes.CreateRoutes()
	go scheduled.StartScheduled()

	select {}
	//reading.ReadInput()

}

func createDataTest() {
	fmt.Println("--------------------START CREATE DATA TEST--------------------\n")

	commands.RunCommand("create database teste_db;")
	commands.RunCommand("use database teste_db;")

	commands.RunCommand("create table teste1 (id integer not_null, nome varchar, idade integer);")
	commands.RunCommand("create unique (unique_teste1_id) table (teste1) columns (id);")

	global.LoadDatabaseInMemory()

	commands.RunCommand("insert into teste1 columns (id, nome, idade) values (1, 'izaque', 26);")
	commands.RunCommand("insert into teste1 columns (id, nome, idade) values (1, 'izaque', 26);")
	commands.RunCommand("insert into teste1 columns (id, nome, idade) values (1, 'izaque', 26);")
	commands.RunCommand("insert into teste1 columns (id, nome, idade) values (1, 'izaque', 26);")

	fmt.Println("--------------------END CREATE DATA TEST--------------------\n")
}
