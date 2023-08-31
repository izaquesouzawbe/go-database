package commands

import (
	"fmt"
	"go-zdb-api/internal/global"
	"go-zdb-api/pkg/file"
)

func commandSelectTable(query string) []string {

	table, _ := extractTable(query)

	lines, _ := file.ReadLines(global.GetPathTableDataRecord(table))
	tamanho := len(lines)

	for _ = range lines {
		/*	if indexes <= 100 {
			fmt.Println(strconv.Itoa(indexes) + " - " + line)
		}*/
	}

	fmt.Println("count: ", tamanho)

	return lines
}
