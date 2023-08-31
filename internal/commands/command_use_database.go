package commands

import "go-zdb-api/internal/global"

func commandUseDatabase(commands []string) {

	databaseName := commands[2]
	global.SetPathDatabase(databaseName)

}
