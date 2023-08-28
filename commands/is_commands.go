package commands

func isCommandCreateDatabase(commands []string) bool {
	return commands[0] == "create" && commands[1] == "database"
}

func isCommandUseDatabase(commands []string) bool {
	return commands[0] == "use" && commands[1] == "database"
}

func isCommandCreateTable(commands []string) bool {
	return commands[0] == "create" && commands[1] == "table"
}

func isCommandAddColumn(commands []string) bool {
	return commands[0] == "add" && commands[1] == "column" && commands[2] == "table"
}

func isCommandInsertInto(commands []string) bool {
	return commands[0] == "insert" && commands[1] == "into"
}

func isCommandSelectTable(commands []string) bool {
	return commands[0] == "select"
}
