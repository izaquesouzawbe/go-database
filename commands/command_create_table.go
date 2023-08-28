package commands

import (
	"go-database/file"
)

func commandCreateTable(query string, commands []string) {

	tableName := commands[2]

	if !file.DirExists(getPathTable(tableName)) {
		file.CreateDir(getPathTable(tableName))
	}

	fields := extractFieldsAndTypesCreateTable(query)

	tableData := Table{
		TableName: tableName,
		Fields:    getFields(fields),
	}

	saveTableConfig(tableData)

}

func getFields(fields map[string]string) []map[string]string {

	newFields := []map[string]string{}

	for field, fieldType := range fields {
		newFields = append(newFields, map[string]string{"name": field, "type": fieldType})
	}

	return newFields

}
