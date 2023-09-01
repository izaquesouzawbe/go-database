package commands_list

import "go-zdb-api/internal/global"

func CommandUseDatabase(commands []string) {

	databaseName := commands[2]
	global.SetPathDatabase(databaseName)

}
