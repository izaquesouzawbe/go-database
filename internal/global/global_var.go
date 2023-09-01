package global

import (
	"fmt"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/dir"
	"go-zdb-api/pkg/file"
)

var TablesInMemory []command.Table
var UniquesInMemory []command.Unique

var pathData string = "data/"
var pathDatabase string

func LoadDatabaseInMemory() {
	loadTablesInMemory()
	loadUniquesInMemory()

	fmt.Println("TablesInMemory: ", TablesInMemory)
	fmt.Println("UniquesInMemory: ", UniquesInMemory)
}

func loadTablesInMemory() {

	tablesDir, _ := dir.ListDir(GetPathTables())
	for _, tableDir := range tablesDir {

		var table command.Table
		file.LoadFileJSON(GetPathTableConfig(tableDir), &table)

		TablesInMemory = append(TablesInMemory, table)
	}

}

func loadUniquesInMemory() {
	tablesDir, _ := dir.ListDir(GetPathTables())
	for _, tableDir := range tablesDir {

		uniquesDir, _ := dir.ListDir(GetPathUniques(tableDir))
		for _, uniqueDir := range uniquesDir {
			var unique command.Unique
			file.LoadFileJSON(GetPathUniqueConfig(tableDir, uniqueDir), &unique)

			UniquesInMemory = append(UniquesInMemory, unique)
		}

	}

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
