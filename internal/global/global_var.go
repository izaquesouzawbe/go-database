package global

import (
	"go-zdb-api/internal/models"
	"go-zdb-api/pkg/file"
)

var TablesInMemory []models.Table
var pathData string = "data/"
var pathDatabase string = "teste_db"

func LoadDatabaseInMemory() {
	loadTablesInMemory()
}

func loadTablesInMemory() {
	var table models.Table
	file.LoadFileJSON(GetPathTableConfig("teste"), &table)

	TablesInMemory = append(TablesInMemory, table)
}

func GetTableInMemory(tableName string) models.Table {

	for _, table := range TablesInMemory {

		if table.TableName == tableName {
			return table
		}
	}

	return models.Table{}
}
