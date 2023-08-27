package commands

import (
	"fmt"
	"go-database/file"
	"go-database/main_aux"
)

var path string = "data/"

func RunCommand(query string) {

	fmt.Println("Query: ", query)
	query = cleanCommand(query)

	commands := getListCommand(query)

	main_aux.RuntimeStarted()

	fmt.Println(commands)

	if isCommandCreateTable(commands) {
		commandCreateTable(commands)
	} else if isCommandAddColumn(commands) {
		commandAddColumn(commands)
	} else if isCommandSelectTable(commands) {
		commandSelectTable(commands)
	} else {
		fmt.Println("Command not found")
	}

	main_aux.RuntimeDone()
}

func commandAddColumn(commands []string) {
	//add column table teste nome text
	columnName := commands[3]
	columnType := commands[4]
	table := commands[5]

	if !file.FileExists(getPathConfigTable(table)) {
		fmt.Println("Table " + table + " not found...")
		return
	}

	line := columnName + ";" + columnType

	err := file.AppendLineToFile(getPathConfigTable(table), line)
	if err != nil {
		fmt.Println("Erro ao adicionar nova linha:", err)
	} else {
		fmt.Println("Nova linha adicionada com sucesso.")
	}

}
