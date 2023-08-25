package main

import (
	"fmt"
	"strconv"
)

func runCommand(command string) {

	fmt.Println("Command: ", command)
	command = cleanCommand(command)
	commands := getListCommand(command)

	runtimeStarted()

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

	runtimeDone()
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

	if !dirExists(getPathTable(table)) {
		createDir(getPathTable(table))
	}

	createFile(getPathTable(table) + "config.txt")

	fmt.Println("Table " + table + " successfully created...")
}

func commandAddColumn(commands []string) {
	//add column table teste nome text
	columnName := commands[3]
	columnType := commands[4]
	table := commands[5]

	if !fileExists(getPathConfigTable(table)) {
		fmt.Println("Table " + table + " not found...")
		return
	}

	line := columnName + ";" + columnType

	err := appendLineToFile(getPathConfigTable(table), line)
	if err != nil {
		fmt.Println("Erro ao adicionar nova linha:", err)
	} else {
		fmt.Println("Nova linha adicionada com sucesso.")
	}

}

func getFields(command string) []string {
	return nil
}

func commandSelectTable(commands []string) {
	//select table teste (id, nome, sobrenome)
	table := commands[2]

	lines, _ := readLines(getPathFilesTable(table)[0])

	for index, line := range lines {
		fmt.Println(strconv.Itoa(index) + " - " + line)
	}

	fmt.Println(table)
}

func whereTable() {

}
