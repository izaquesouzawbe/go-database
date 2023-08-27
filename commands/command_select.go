package commands

import (
	"fmt"
	"go-database/file"
)

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
