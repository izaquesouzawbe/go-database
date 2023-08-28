package commands

import (
	"fmt"
	"go-database/file"
	"go-database/main_aux"
	"log"
)

func RunCommand(query string) []string {

	main_aux.RuntimeStarted()

	log.Println("Query: ", query)

	query = cleanQuery(query)
	commands := getListCommand(query)

	var stringReturn []string

	switch {
	case isCommandCreateDatabase(commands):
		commandCreateDatabase(commands)

	case isCommandUseDatabase(commands):
		commandUseDatabase(commands)

	case isCommandCreateTable(commands):
		commandCreateTable(query, commands)

	case isCommandAddColumn(commands):
		commandAddColumn(commands)

	case isCommandInsertInto(commands):
		commandInsertInto(query)

	case isCommandSelectTable(commands):
		stringReturn = commandSelectTable(query)

	default:
		fmt.Println("Command not found")
	}

	main_aux.RuntimeDone()

	return stringReturn

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
