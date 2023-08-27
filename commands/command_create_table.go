package commands

import (
	"fmt"
	"go-database/file"
)

func commandCreateTable(commands []string) {

	table := commands[2]

	if !file.DirExists(getPathTable(table)) {
		file.CreateDir(getPathTable(table))
	}

	file.CreateFile(getPathTable(table) + "config.txt")

	fmt.Println("Table " + table + " successfully created...")
}
