package commands

import (
	"fmt"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/general"
)

var table command.Table

func RunCommand(command string) []map[string]string {

	general.RuntimeStarted()

	command = cleanCommand(command)
	querys := getQuerys(command)

	/*	if batch {
		commandInsertIntoQuerys(querys)
		general.RuntimeDone()
		return []string{}
	}*/

	for _, query := range querys {

		commands := getCommands(query)

		fmt.Println(commands)

		switch {
		case isCommandCreateDatabase(commands):
			commandCreateDatabase(commands)

		case isCommandUseDatabase(commands):
			commandUseDatabase(commands)

		case isCommandCreateTable(commands):
			commandCreateTable(query, commands)

		case isCommandCreateSequence(commands):
			commandCreateSequence(commands)

		case isCommandCreateUnique(commands):
			commandCreateUnique(query)

		case isCommandInsertInto(commands):
			commandInsertInto(query)

		case isCommandSelectTable(commands):
			return commandSelectTable(query)

		default:
			fmt.Println("Command not found")
		}
	}

	general.RuntimeDone()
	return nil

}

func commandAddColumn(commands []string) {

}
