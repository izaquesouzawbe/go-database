package commands

import (
	"fmt"
	"go-zdb-api/pkg/file"
	"go-zdb-api/pkg/main_aux"
)

var table Table

func RunCommand(query string, batch bool) []string {

	err := file.LoadFileJSON(getPathConfigTable("teste"), &table)
	if err != nil {
		return nil
	}

	main_aux.RuntimeStarted()

	query = cleanQuery(query)
	querys := getQuerys(query)

	//log.Println("Querys: ", querys)

	if batch {
		commandInsertIntoQuerys(querys)
		main_aux.RuntimeDone()
		return []string{}
	}

	var stringReturn []string

	for _, query := range querys {

		commands := getCommands(query)

		switch {
		case isCommandCreateDatabase(commands):
			commandCreateDatabase(commands)

		case isCommandUseDatabase(commands):
			commandUseDatabase(commands)

		case isCommandCreateTable(commands):
			stringReturn = commandCreateTable(query, commands)

		case isCommandCreateSequence(commands):
			stringReturn = commandCreateSequence(commands)

		/*case isCommandInsertInto(commands):
		stringReturn = commandInsertInto(querys)*/

		case isCommandSelectTable(commands):
			stringReturn = commandSelectTable(query)

		default:
			fmt.Println("Command not found")
		}
	}

	main_aux.RuntimeDone()

	return stringReturn

}

func commandAddColumn(commands []string) {

}
