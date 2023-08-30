package commands

import (
	"fmt"
	"go-zdb-api/pkg/file"
)

func commandSelectTable(query string) []string {

	table, _ := extractTable(query)

	lines, _ := file.ReadLines(getPathTableDataRecord(table))
	tamanho := len(lines)

	for _ = range lines {
		/*	if indexes <= 100 {
			fmt.Println(strconv.Itoa(indexes) + " - " + line)
		}*/
	}

	fmt.Println("count: ", tamanho)

	return lines
}
