package commands

import (
	"encoding/json"
	"go-database/file"
	"log"
)

type Table struct {
	TableName string              `json:"table_name"`
	Fields    []map[string]string `json:"fields"`
}

func commandCreateTable(query string, commands []string) {

	tableName := commands[2]

	if !file.DirExists(getPathTable(tableName)) {
		file.CreateDir(getPathTable(tableName))
	}

	fields := extractFieldsAndTypes(query)

	for field, fieldType := range fields {
		log.Printf("%s: %s\n", field, fieldType)
	}

	tableData := Table{
		TableName: tableName,
		Fields:    getFields(fields),
	}

	file := file.CreateFile(getPathConfigTable(tableName), true)
	encoder := json.NewEncoder(file)
	encoder.Encode(tableData)
}

func getFields(map[string]string) []map[string]string {

	return []map[string]string{
		{"name": "id", "type": "int"},
		{"name": "nome", "type": "varchar"},
		{"name": "email", "type": "varchar"},
		{"name": "data_nascimento", "type": "date"},
	}

}
