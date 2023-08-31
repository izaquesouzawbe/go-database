package global

import (
	"fmt"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
)

var TablesInMemory []command.Table
var UniquesInMemory []command.Unique

var pathData string = "data/"
var pathDatabase string = "teste_db"

func LoadDatabaseInMemory() {
	loadTablesInMemory()
	loadUniquesInMemory()
}

func loadTablesInMemory() {
	var table command.Table
	file.LoadFileJSON(GetPathTableConfig("teste"), &table)

	TablesInMemory = append(TablesInMemory, table)
}

func loadUniquesInMemory() {
	var unique command.Unique
	file.LoadFileJSON(GetPathUniqueConfig("simple", "unique_simple_symbol"), &unique)

	UniquesInMemory = append(UniquesInMemory, unique)
}

func GetTableInMemory(tableName string) command.Table {

	for _, table := range TablesInMemory {

		if table.TableName == tableName {
			return table
		}
	}

	return command.Table{}
}

func GetUniqueInMemory(tableName string) command.Unique {

	fmt.Println(UniquesInMemory)

	for _, unique := range UniquesInMemory {

		if unique.Table == tableName {
			return unique
		}
	}

	return command.Unique{}
}
