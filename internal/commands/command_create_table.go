package commands

import (
	"go-zdb-api/internal/models"
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

	tableData := models.Table{
		TableName: tableName,
		Fields:    fields,
	}

	saveTableConfig(tableData)

	return []string{"Table '" + tableName + "' successfully created!"}
}

func createDirTable(tableName string) {

	if !file.DirExists(getPathTable(tableName)) {
		file.CreateDir(getPathTable(tableName))
	}

	if !file.DirExists(getPathTableData(tableName)) {
		file.CreateDir(getPathTableData(tableName))
	}
}

func onValidateCreateTable(query string) string {

	return ""

}
