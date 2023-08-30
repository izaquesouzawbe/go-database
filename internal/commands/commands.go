package commands

import (
	"fmt"
	"go-zdb-api/internal/models"
	"go-zdb-api/pkg/file"
	"go-zdb-api/pkg/general"
)

var table models.Table

func RunCommand(query string, batch bool) []string {

	err := file.LoadFileJSON(getPathConfigTable("teste"), &table)
	if err != nil {
		return nil
	}

	general.RuntimeStarted()

	query = cleanQuery(query)
	querys := getQuerys(query)

	//log.Println("Querys: ", querys)

	if batch {
		commandInsertIntoQuerys(querys)
		general.RuntimeDone()
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

	general.RuntimeDone()

	return stringReturn

}

func commandAddColumn(commands []string) {

}
