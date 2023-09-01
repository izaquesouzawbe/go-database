package commands_list

import (
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/global"
	"go-zdb-api/pkg/file"
	"strings"
)

func CommandSelectTable(query string) []map[string]string {
	tableName, _ := commands_func.ExtractTableSelect(query)
	wheres := commands_func.ExtractKeyValueWhere(query)

	lines, _ := file.ReadLines(global.GetPathTableDataRecord(tableName))

	table := global.GetTableInMemory(tableName)

	var records []map[string]string

	for _, line := range lines {

		if len(line) > 0 {

			lineFree := false

			linesValue := strings.Split(line, ";")
			record := make(map[string]string)

			for i, field := range table.Fields {
				record[field.Name] = linesValue[i]

				//WHERE
				if len(wheres) > 0 {

					for whereField, whereValue := range wheres {

						if whereField == field.Name && strings.ToLower(linesValue[i]) == whereValue {
							lineFree = true
						}
					}
				} else {
					lineFree = true
				}
			}

			if lineFree {
				records = append(records, record)
			}
		}

	}

	return records
}
