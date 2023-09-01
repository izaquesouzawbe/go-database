package commands_list

import (
	"fmt"
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/internal/models/request"
	"go-zdb-api/pkg/file"
	"log"
	"reflect"
	"strconv"
)

type InsertLine struct {
	LinesRecord LineRecord
	LinesUnique LineUnique
}

type LineRecord struct {
	Table string
	Lines []string
}

type LineUnique struct {
	Unique string
	Lines  []string
}

func CommandInsertInto(query string) string {

	msg := onValidateInsertInto(query)
	if len(msg) > 0 {
		return ""
	}

	tableName := commands_func.ExtractInsertTable(query)
	insertRequest := request.InsertRequest{
		Table:   tableName,
		Records: getRecordsQuery(query),
	}

	table := global.GetTableInMemory(insertRequest.Table)

	msg = onValidateNotNullInsertInto(table.Fields, insertRequest)
	if len(msg) > 0 {
		return ""
	}

	lines := getLines(table, insertRequest)

	saveInsertLines(lines)

	return ""
}

func saveInsertLines(insertLine InsertLine) {

	if len(insertLine.LinesRecord.Lines) > 0 {
		file.AppendLineToFile(global.GetPathTableDataRecord(insertLine.LinesRecord.Table), insertLine.LinesRecord.Lines)
	}

	if len(insertLine.LinesUnique.Lines) > 0 {
		file.AppendLineToFile(global.GetPathUniqueData(insertLine.LinesRecord.Table, insertLine.LinesUnique.Unique), insertLine.LinesUnique.Lines)
	}
}

func RouteInsertInto(insertRequest request.InsertRequest) string {

	table := global.GetTableInMemory(insertRequest.Table)
	msg := onValidateNotNullInsertInto(table.Fields, insertRequest)
	if len(msg) > 0 {
		return ""
	}

	lines := getLines(table, insertRequest)

	saveInsertLines(lines)

	return ""

}

func getLines(table command.Table, insertRequest request.InsertRequest) InsertLine {

	countLines := commands_func.CountLines(global.GetPathTableDataRecord(insertRequest.Table))
	fmt.Println("countLines ", countLines)

	unique := global.GetUniqueInMemory(insertRequest.Table)
	processUnique := !reflect.ValueOf(unique).IsZero()

	var lines []string
	var linesUnique []string

	for _, record := range insertRequest.Records {

		line := getLineValues(table.Fields, record)

		if processUnique {

			lineUnique := ""

			countLines++
			/*for _, fieldValue := range {

			}*/
			lineUnique += strconv.Itoa(countLines)
			linesUnique = append(linesUnique, lineUnique)

		}

		lines = append(lines, line)
	}

	insertLine := InsertLine{
		LinesRecord: LineRecord{
			Table: table.TableName,
			Lines: lines,
		},
		LinesUnique: LineUnique{
			Unique: unique.Name,
			Lines:  linesUnique,
		},
	}

	return insertLine
}

func getLineValues(fields []command.Field, record map[string]string) string {

	line := ""
	for _, field := range fields {

		lineValue := ""

		for fieldName, fieldValue := range record {

			if fieldName == field.Name {
				lineValue += fieldValue + ";"
			}
		}

		if len(lineValue) > 0 {
			line += lineValue
		} else {
			line += ";"
		}
	}
	return line
}

func getRecordsQuery(query string) []map[string]string {
	var records []map[string]string
	record := make(map[string]string)

	columns := commands_func.ExtractInsertColumns(query)
	values := commands_func.ExtractInsertValues(query)

	for i, column := range columns {
		record[column] = values[i]
	}

	records = append(records, record)

	return records
}

func onValidateInsertInto(query string) string {
	return ""
}

func onValidateNotNullInsertInto(fields []command.Field, insertCommand request.InsertRequest) string {

	for _, field := range fields {

		if field.NotNull == 1 {
			for columnName, columnValue := range insertCommand.Records {
				log.Println(columnName, columnValue)
				/*if field.Name == columnName && len(columnValue) > 0 {
					return ""
				} else {
					return "Error: Field '" + field.Name + "' cannot be null"
				}*/
			}
		}
	}

	return ""
}
