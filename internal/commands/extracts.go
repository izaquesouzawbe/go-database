package commands

import (
	"fmt"
	"go-zdb-api/internal/models/command"
	"regexp"
	"strings"
)

func extractUniqueName(query string) string {

	padrao := `unique\s*\((.*?)\)`
	r := regexp.MustCompile(padrao)
	correspondencias := r.FindStringSubmatch(query)

	if len(correspondencias) > 1 {
		return correspondencias[1]
	} else {
		return ""
	}
}

func extractUniqueTable(query string) string {
	padrao := `table\s*\((.*?)\)`
	r := regexp.MustCompile(padrao)
	correspondencias := r.FindStringSubmatch(query)

	if len(correspondencias) > 1 {
		return correspondencias[1]
	} else {
		return ""
	}
}

func extractUniqueColumns(query string) []string {
	padrao := `column\s*\((.*?)\)`
	r := regexp.MustCompile(padrao)
	correspondencias := r.FindStringSubmatch(query)

	if len(correspondencias) > 1 {
		colunas := strings.Split(correspondencias[1], ", ")
		return colunas
	} else {
		return nil
	}
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

	re := regexp.MustCompile(`where\s+(.*?)$`)

	matches := re.FindStringSubmatch(query)

	if len(matches) >= 2 {
		conditions := strings.TrimSpace(matches[1])
		pairs := strings.Split(conditions, " and ")
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
