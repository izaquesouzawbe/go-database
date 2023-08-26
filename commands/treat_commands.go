package commands

import (
	"regexp"
	"strings"
)

func cleanCommand(command string) string {

	text := strings.ToLower(command)

	text = strings.TrimSpace(text)
	text = reduceToSingleSpace(text)

	return text
}

func getListCommand(command string) []string {

	stringList := strings.Split(command, " ")
	return stringList
}

func reduceToSingleSpace(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}
