package commands_func

func IsCommandCreateDatabase(commands []string) bool {
	return commands[0] == "create" && commands[1] == "database"
}

func IsCommandUseDatabase(commands []string) bool {
	return commands[0] == "use" && commands[1] == "database"
}

func IsCommandCreateTable(commands []string) bool {
	return commands[0] == "create" && commands[1] == "table"
}

func IsCommandCreateSequence(commands []string) bool {
	return commands[0] == "create" && commands[1] == "sequence"
}

func IsCommandAddColumn(commands []string) bool {
	return commands[0] == "add" && commands[1] == "column" && commands[2] == "tables"
}

func IsCommandInsertInto(commands []string) bool {
	return commands[0] == "insert" && commands[1] == "into"
}

func IsCommandSelectTable(commands []string) bool {
	return commands[0] == "select"
}

func IsCommandCreateUnique(commands []string) bool {
	return commands[0] == "create" && commands[1] == "unique"
}
