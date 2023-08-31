package global

import (
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
)

var TablesInMemory []command.Table
var pathData string = "data/"
var pathDatabase string = "teste_db"

func LoadDatabaseInMemory() {
	loadTablesInMemory()
}

func loadTablesInMemory() {
	var table command.Table
	file.LoadFileJSON(GetPathTableConfig("teste"), &table)

	TablesInMemory = append(TablesInMemory, table)
}

func GetTableInMemory(tableName string) command.Table {

	for _, table := range TablesInMemory {

		if table.TableName == tableName {
			return table
		}
	}

	return command.Table{}
}
