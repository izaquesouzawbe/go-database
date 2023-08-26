package commands

import (
	"fmt"
	"go-database/file"
	"go-database/main_aux"
	"regexp"
	"strings"
)

var path string = "data/"

func RunCommand(command string) {

	fmt.Println("Command: ", command)
	command = cleanCommand(command)
	commands := getListCommand(command)

	main_aux.RuntimeStarted()

	fmt.Println(commands)

	if isCommandCreateTable(commands) {
		commandCreateTable(commands)
	} else if isCommandAddColumn(commands) {
		commandAddColumn(commands)
	} else if isCommandSelectTable(commands) {
		commandSelectTable(commands)
	} else {
		fmt.Println("Command not found")
	}

	main_aux.RuntimeDone()
}

func getPathTable(table string) string {
	return path + table + "/"
}

func getPathConfigTable(table string) string {
	return getPathTable(table) + "config.txt"
}

func getPathFilesTable(table string) []string {

	newPath := getPathTable(table)

	var files []string
	files = append(files, newPath+"1.txt")
	return files
}

func commandCreateTable(commands []string) {

	table := commands[2]

	if !file.DirExists(getPathTable(table)) {
		file.CreateDir(getPathTable(table))
	}

	file.CreateFile(getPathTable(table) + "config.txt")

	fmt.Println("Table " + table + " successfully created...")
}

func commandAddColumn(commands []string) {
	//add column table teste nome text
	columnName := commands[3]
	columnType := commands[4]
	table := commands[5]

	if !file.FileExists(getPathConfigTable(table)) {
		fmt.Println("Table " + table + " not found...")
		return
	}

	line := columnName + ";" + columnType

	err := file.AppendLineToFile(getPathConfigTable(table), line)
	if err != nil {
		fmt.Println("Erro ao adicionar nova linha:", err)
	} else {
		fmt.Println("Nova linha adicionada com sucesso.")
	}

}

func commandSelectTable(commands []string) []string {
	//select * from teste
	table := commands[2]

	lines, _ := file.ReadLines(getPathFilesTable(table)[0])

	/*for index, line := range lines {
		if index <= 100 {
			fmt.Println(strconv.Itoa(index) + " - " + line)
		}
	}*/

	fmt.Println(table)

	return lines
}

func extractColumns(query string) []string {

	query = cleanCommand(query)

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
