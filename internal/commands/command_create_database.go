package commands

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/pkg/file"
)

func commandCreateDatabase(commands []string) {

	databaseName := commands[2]

	global.SetPathDatabase(databaseName)

	file.CreateDir(global.GetPathDatabase())
	file.CreateDir(global.GetPathTables())
	file.CreateDir(global.GetPathSequences())
	file.CreateDir(global.GetPathIndexes())

}
