package commands

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models"
	"go-zdb-api/pkg/file"
)

func commandInsertInto(query string) string {

	msg := onValidateInsertInto(query)
	if len(msg) > 0 {
		return ""
	}

	insertCommand, _ := extractInsertCommand(query)
	table := global.GetTableInMemory(insertCommand.TableName)

	insertCommandValues, _ := extractInsertCommandValues(query)

	msg = onValidateNotNullInsertInto(table.Fields, insertCommand, insertCommandValues)
	if len(msg) > 0 {
		return ""
	}

	line := ""
	for _, field := range table.Fields {
		line += getFieldValue(field, insertCommand, insertCommandValues)
	}

	file.AppendLineToFile(global.GetPathTableDataRecord(insertCommand.TableName), []string{line})

	return line
}

func commandInsertIntoQuerys(querys []string) []string {

	var lines []string

	for _, query := range querys {
		if len(query) > 0 {
			lines = append(lines, commandInsertInto(query))
		}
	}

	return []string{}
}

func getFieldValue(field models.Field, insertCommand *models.InsertCommand, insertCommandValues *models.InsertCommandValues) string {

	for iColumn, column := range insertCommand.Fields {

		if column == field.Name {
			return insertCommandValues.Values[iColumn] + ";"
		}

	}

	return ";"
}

func onValidateInsertInto(query string) string {
	return ""
}

func onValidateNotNullInsertInto(fields []models.Field, insertCommand *models.InsertCommand, insertCommandValues *models.InsertCommandValues) string {

	for _, field := range fields {

		if field.NotNull == 1 {
			for i, column := range insertCommand.Fields {
				if field.Name == column && len(insertCommandValues.Values[i]) > 0 {
					return ""
				} else {
					return "Error: Field '" + field.Name + "' cannot be null"
				}
			}
		}
	}

	return ""
}
