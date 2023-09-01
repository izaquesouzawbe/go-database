package commands_func

import (
	"go-zdb-api/pkg/dir"
	"regexp"
	"strings"
)

func CreateDirData() {
	if !dir.DirExists("data/") {
		dir.CreateDir("data/")
	}
}

func CleanCommand(query string) string {

	query = strings.ToLower(query)
	query = strings.TrimSpace(query)
	query = reduceToSingleSpace(query)

	return query
}

func GetCommands(command string) []string {

	commands := strings.Split(command, " ")
	return commands
}

func GetQuerys(query string) []string {
	querys := strings.Split(query, ";")
	return querys
}

func reduceToSingleSpace(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}

func getSequenceOfFields(fieldsMap []map[string]string) string {

	for _, field := range fieldsMap {
		if field["type"] == "sequence" {
			return field["name"]
		}
	}

	return ""
}

func findIndexFields(list []string, value string) int {
	for i, valor := range list {
		if valor == value {
			return i
		}
	}
	return -1
}
