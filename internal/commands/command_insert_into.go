package commands

import (
	"go-zdb-api/internal/models"
	"go-zdb-api/pkg/file"
)

func commandInsertInto(query string, insertCommand *models.InsertCommand) string {

	msg := onValidateInsertInto(query)
	if len(msg) > 0 {
		return ""
	}

	//insertCommand, _ := extractInsertCommandInfo(query)

	msg = onValidateNotNullInsertInto(table.Fields, insertCommand)
	if len(msg) > 0 {
		return ""
	}

	line := ""

	for _, field := range table.Fields {
		line += getFieldValue(field, insertCommand)
	}

	return line
}

func commandInsertIntoQuerys(querys []string) []string {

	insertCommand, _ := extractInsertCommandInfo(querys[0])

	var lines []string

	for _, query := range querys {
		if len(query) > 0 {
			lines = append(lines, commandInsertInto(query, insertCommand))
		}
	}

	file.AppendLineToFile(getPathTableDataRecord(insertCommand.TableName), lines)

	return []string{}
}

func getFieldValue(field models.Field, insertCommand *models.InsertCommand) string {

	for iColumn, column := range insertCommand.Fields {

		if column == field.Name {
			return insertCommand.Values[iColumn] + ";"
		}

	}

	return ";"
}

func onValidateInsertInto(query string) string {
	return ""
}

func onValidateNotNullInsertInto(fields []models.Field, insertCommand *models.InsertCommand) string {

	for _, field := range fields {

		if field.NotNull == 1 {
			for i, column := range insertCommand.Fields {
				if field.Name == column && len(insertCommand.Values[i]) > 0 {
					return ""
				} else {
					return "Error: Field '" + field.Name + "' cannot be null"
				}
			}
		}
	}

	return ""
}
