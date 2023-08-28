package commands

import (
	"encoding/json"
	"fmt"
	"go-database/file"
	"regexp"
	"strings"
)

func CreateDirData() {
	if !file.DirExists("data/") {
		file.CreateDir("data/")
	}
}

func cleanQuery(query string) string {

	query = strings.ToLower(query)
	query = strings.TrimSpace(query)
	query = reduceToSingleSpace(query)

	return query
}

func getListCommand(command string) []string {

	stringList := strings.Split(command, " ")
	return stringList
}

func reduceToSingleSpace(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}

func saveTableConfig(tableData Table) {

	file := file.CreateFile(getPathConfigTable(tableData.TableName), true)
	encoder := json.NewEncoder(file)
	encoder.Encode(tableData)

}

func saveDatabaseConfig(databaseData Database) {

	file := file.CreateFile(getPathConfigDatabase(), true)
	encoder := json.NewEncoder(file)
	encoder.Encode(databaseData)

}

func extractTable(query string) (string, error) {
	re := regexp.MustCompile(`from\s+([^\s;]+)`)
	matches := re.FindStringSubmatch(query)

	if len(matches) < 2 {
		return "", fmt.Errorf("Não foi possível extrair a tabela")
	}

	return matches[1], nil
}

func extractColumnsSelect(query string) []string {

	re := regexp.MustCompile(`SELECT\s+(.*?)\s+FROM`)
	matches := re.FindStringSubmatch(query)
	if len(matches) >= 2 {
		columns := strings.Split(matches[1], ",")
		for i, col := range columns {
			columns[i] = strings.TrimSpace(col)
		}
		return columns
	}
	return nil

}

func extractKeyValueWhere(query string) map[string]string {
	re := regexp.MustCompile(`WHERE\s+(.*?)$`)
	matches := re.FindStringSubmatch(query)
	if len(matches) >= 2 {
		conditions := strings.TrimSpace(matches[1])
		pairs := strings.Split(conditions, " AND ")
		keyValuePairs := make(map[string]string)

		for _, pair := range pairs {
			parts := strings.Split(pair, "=")
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				// Remove single quotes around the value
				value = strings.Trim(value, "'")
				keyValuePairs[key] = value
			}
		}

		return keyValuePairs
	}
	return nil
}

func extractFieldsAndTypesCreateTable(input string) map[string]string {
	result := make(map[string]string)
	re := regexp.MustCompile(`\((.*?)\)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 2 {
		fields := strings.Split(matches[1], ",")
		for _, field := range fields {
			parts := strings.Split(strings.TrimSpace(field), " ")
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			}
		}
	}

	return result
}

type InsertCommand struct {
	TableName string
	Fields    []string
	Values    []string
}

func extractInsertCommandInfo(texto string) (*InsertCommand, error) {
	re := regexp.MustCompile(`insert into ([^\s(]+)\(([^)]+)\) values \(([^)]+)\)`)
	matches := re.FindStringSubmatch(texto)

	if len(matches) < 4 {
		return nil, fmt.Errorf("Não foi possível fazer o parsing do texto")
	}

	tabela := matches[1]
	campos := matches[2]
	valores := matches[3]

	camposSeparados := regexp.MustCompile(`\s*,\s*`).Split(campos, -1)
	valoresSeparados := regexp.MustCompile(`\s*,\s*`).Split(valores, -1)

	return &InsertCommand{
		TableName: tabela,
		Fields:    camposSeparados,
		Values:    valoresSeparados,
	}, nil
}
