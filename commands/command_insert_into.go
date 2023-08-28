package commands

import (
	"fmt"
	"go-database/file"
)

func commandInsertInto(query string) {

	insertCommand, _ := extractInsertCommandInfo(query)
	var tabelaJson Table

	err := file.LoadFileJSON(getPathConfigTable(insertCommand.TableName), &tabelaJson)
	if err != nil {
		return
	}

	line := ""

	for iField, field := range tabelaJson.Fields {

		fmt.Println(iField)
		fmt.Println(field["name"])

		for iColumn, column := range insertCommand.Fields {

			if column == field["name"] {
				line += insertCommand.Values[iColumn] + ";"
			}

		}

	}

	err = file.AppendLineToFile(getPathDataTable(insertCommand.TableName), line)
	if err != nil {
		return
	}

}
