package commands

import (
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/commands/commands_list"
	"go-zdb-api/pkg/general"
	"log"
)

func RunCommand(command string) []map[string]string {

	general.RuntimeStarted()

	command = commands_func.CleanCommand(command)
	querys := commands_func.GetQuerys(command)

	for _, query := range querys {

		if len(query) == 0 {
			continue
		}

		log.Println("Query: " + query)

		commands := commands_func.GetCommands(query)

		switch {
		case commands_func.IsCommandCreateDatabase(commands):
			commands_list.CommandCreateDatabase(commands)

		case commands_func.IsCommandUseDatabase(commands):
			commands_list.CommandUseDatabase(commands)

		case commands_func.IsCommandCreateTable(commands):
			commands_list.CommandCreateTable(query, commands)

		case commands_func.IsCommandCreateSequence(commands):
			commands_list.CommandCreateSequence(commands)

		case commands_func.IsCommandCreateUnique(commands):
			commands_list.CommandCreateUnique(query)

		case commands_func.IsCommandInsertInto(commands):
			commands_list.CommandInsertInto(query)

		case commands_func.IsCommandSelectTable(commands):
			return commands_list.CommandSelectTable(query)

		default:
			log.Println("Command not found")
		}
	}

	general.RuntimeDone()
	return nil

}

func commandAddColumn(commands []string) {

}
