package commands

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
)

func commandCreateTable(query string, commands []string) []string {

	msg := onValidateCreateTable(query)
	if len(msg) > 0 {
		return []string{msg}
	}

	tableName := commands[2]

	createDirTable(tableName)
	fields := extractFieldsCreateTable(query)

	tableData := command.Table{
		TableName: tableName,
		Fields:    fields,
	}

	saveTableConfig(tableData)

	return []string{"Table '" + tableName + "' successfully created!"}
}

func createDirTable(tableName string) {

	if !file.DirExists(global.GetPathTable(tableName)) {
		file.CreateDir(global.GetPathTable(tableName))
	}

	if !file.DirExists(global.GetPathTableData(tableName)) {
		file.CreateDir(global.GetPathTableData(tableName))
	}
}

func onValidateCreateTable(query string) string {

	return ""

}
