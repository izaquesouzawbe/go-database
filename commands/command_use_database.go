package commands

func commandUseDatabase(commands []string) {

	databaseName := commands[2]
	setPathDatabase(databaseName)

}
