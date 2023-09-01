package commands_list

import (
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/dir"
)

func CommandCreateTable(query string, commands []string) []string {

	msg := onValidateCreateTable(query)
	if len(msg) > 0 {
		return []string{msg}
	}

	tableName := commands[2]

	createDirTable(tableName)
	fields := commands_func.ExtractFieldsCreateTable(query)

	tableData := command.Table{
		TableName: tableName,
		Fields:    fields,
	}

	commands_func.SaveTableConfig(tableData)

	return []string{"Table '" + tableName + "' successfully created!"}
}

func createDirTable(tableName string) {

	if !dir.DirExists(global.GetPathTable(tableName)) {
		dir.CreateDir(global.GetPathTable(tableName))
	}

	if !dir.DirExists(global.GetPathTableData(tableName)) {
		dir.CreateDir(global.GetPathTableData(tableName))
	}
}

func onValidateCreateTable(query string) string {

	return ""

}
