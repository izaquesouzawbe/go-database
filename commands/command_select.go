package commands

import (
	"fmt"
	"go-database/file"
)

func commandSelectTable(query string) []string {

	table, _ := extractTable(query)

	lines, _ := file.ReadLines(getPathDataTable(table))
	tamanho := len(lines)

	for _ = range lines {
		/*	if index <= 100 {
			fmt.Println(strconv.Itoa(index) + " - " + line)
		}*/
	}

	fmt.Println("count: ", tamanho)

	return lines
}
