package commands

import (
	"encoding/json"
	"go-database/file"
	"regexp"
	"strings"
)

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

func extractColumns(query string) []string {

	query = cleanQuery(query)

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

func extractKeyValuePairs(query string) map[string]string {
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

func extractFieldsAndTypes(input string) map[string]string {
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
