package commands

func isCommandCreateTable(commands []string) bool {
	//create table table_name
	return commands[0] == "create" && commands[1] == "table"
}

func isCommandAddColumn(commands []string) bool {
	//add column table teste nome text
	return commands[0] == "add" && commands[1] == "column" && commands[2] == "table"
}

func isCommandSelectTable(commands []string) bool {
	//select table teste (id, nome, sobrenome)
	return commands[0] == "select" && commands[1] == "table"
}
