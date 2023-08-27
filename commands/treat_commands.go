package commands

import (
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
