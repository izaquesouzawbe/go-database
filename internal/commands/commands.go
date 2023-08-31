package commands

import (
	"fmt"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/general"
)

var table command.Table

func RunCommand(command string) []string {

	general.RuntimeStarted()

	command = cleanCommand(command)
	querys := getQuerys(command)

	/*	if batch {
		commandInsertIntoQuerys(querys)
		general.RuntimeDone()
		return []string{}
	}*/

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

		case isCommandInsertInto(commands):
			stringReturn = []string{commandInsertInto(query)}

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
