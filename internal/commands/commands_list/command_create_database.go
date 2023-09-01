package commands_list

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/pkg/dir"
	"log"
)

func CommandCreateDatabase(commands []string) {

	databaseName := commands[2]

	global.SetPathDatabase(databaseName)

	verifyCreateDir(global.GetPathDatabase(), "Database '"+databaseName+"' has already been created")
	verifyCreateDir(global.GetPathTables(), "")
	verifyCreateDir(global.GetPathSequences(), "")

}

func verifyCreateDir(path string, msgError string) {
	if !dir.DirExists(path) {
		dir.CreateDir(path)
	} else {
		if len(msgError) > 0 {
			log.Println(msgError)
		}
	}
}
