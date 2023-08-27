package commands

import "go-database/file"

func commandCreateDatabase(commands []string) {

	databaseName := commands[2]
	file.CreateDir(getPathData() + databaseName)

}
