package commands

import (
	"encoding/json"
	"fmt"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
	"regexp"
	"strings"
)

func CreateDirData() {
	if !file.DirExists("data/") {
		file.CreateDir("data/")
	}
}

func cleanCommand(query string) string {

	query = strings.ToLower(query)
	query = strings.TrimSpace(query)
	query = reduceToSingleSpace(query)

	return query
}

func getCommands(command string) []string {

	commands := strings.Split(command, " ")
	return commands
}

func getQuerys(query string) []string {
	querys := strings.Split(query, ";")
	return querys
}

func reduceToSingleSpace(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}

func saveTableConfig(tableData command.Table) {

	file := file.CreateFile(global.GetPathTableConfig(tableData.TableName), true)
	encoder := json.NewEncoder(file)
	encoder.Encode(tableData)

}

func getSequenceOfFields(fieldsMap []map[string]string) string {

	for _, field := range fieldsMap {
		if field["type"] == "sequence" {
			return field["name"]
		}
	}

	return ""
}

func saveSequenceConfig(sequence command.Sequence) {
	file := file.CreateFile(global.GetPathSequence(sequence.Name), true)

	encoder := json.NewEncoder(file)
	encoder.Encode(sequence)
}

func saveDatabaseConfig(databaseData command.Database) {

	file := file.CreateFile(global.GetPathConfigDatabase(), true)
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

func findIndexFields(list []string, value string) int {
	for i, valor := range list {
		if valor == value {
			return i
		}
	}
	return -1
}

func extractFieldsCreateTable(input string) []command.Field {

	var result []command.Field

	re := regexp.MustCompile(`\((.*?)\)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 2 {

		fields := strings.Split(matches[1], ",")

		for _, field := range fields {
			parts := strings.Split(strings.TrimSpace(field), " ")

			if len(parts) >= 2 {

				field := command.Field{
					Name:    parts[0],
					Type:    parts[1],
					NotNull: 0,
				}

				var index int = 0

				index = findIndexFields(parts, "not_null")
				if index > 0 {
					field.NotNull = 1
				}

				result = append(result, field)
			}
		}
	}

	return result
}

func extractInsertCommand(texto string) (*command.InsertCommand, error) {
	re := regexp.MustCompile(`insert into ([^\s(]+)\(([^)]+)\)`)
	matches := re.FindStringSubmatch(texto)

	if len(matches) < 2 {
		return nil, fmt.Errorf("Não foi possível fazer o parsing do texto")
	}

	tabela := matches[1]
	campos := matches[2]

	camposSeparados := regexp.MustCompile(`\s*,\s*`).Split(campos, -1)

	return &command.InsertCommand{
		TableName: tabela,
		Fields:    camposSeparados,
	}, nil
}

func extractInsertCommandValues(texto string) (*command.InsertCommandValues, error) {

	re := regexp.MustCompile(`values \(([^)]+)\)`)
	matches := re.FindStringSubmatch(texto)

	if len(matches) < 1 {
		return nil, fmt.Errorf("Não foi possível fazer o parsing do texto")
	}

	values := matches[1]

	valoresSeparados := regexp.MustCompile(`\s*,\s*`).Split(values, -1)

	return &command.InsertCommandValues{
		Values: valoresSeparados,
	}, nil
}
