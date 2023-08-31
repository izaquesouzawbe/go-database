package commands

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/pkg/file"
	"strings"
)

func commandSelectTable(query string) []map[string]string {
	tableName, _ := extractTable(query)
	wheres := extractKeyValueWhere(query)

	lines, _ := file.ReadLines(global.GetPathTableDataRecord(tableName))

	file.LoadFileJSON(global.GetPathTableConfig(tableName), &table)

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
