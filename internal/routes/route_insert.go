package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/commands"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/internal/models/request"
	"go-zdb-api/pkg/file"
	"go-zdb-api/pkg/general"
	"reflect"
	"strconv"
)

func routeInsert(c *gin.Context) {

	general.RuntimeStarted()

	var insert request.InsertRequest
	err := c.BindJSON(&insert)
	if err != nil {
		return
	}

	var table command.Table
	file.LoadFileJSON(global.GetPathTableConfig(insert.Table), &table)

	var lines []string
	var linesIndex []string

	countLines := commands.CountLines(global.GetPathTableDataRecord(insert.Table))
	fmt.Println("countLines ", countLines)

	unique := global.GetUniqueInMemory(insert.Table)
	processUnique := !reflect.ValueOf(unique).IsZero()

	for _, record := range insert.Records {

		line := ""

		for _, field := range table.Fields {

			lineValue := ""

			for fieldName, fieldValue := range record {

				if fieldName == field.Name {
					lineValue += fieldValue + ";"
				}
			}

			if len(lineValue) > 0 {
				line += lineValue
			} else {
				line += ";"
			}
		}

		if processUnique {

			lineIndex := ""

			countLines++
			/*			for _, fieldValue := range {

						}*/
			lineIndex += strconv.Itoa(countLines)
			linesIndex = append(linesIndex, lineIndex)

		}

		lines = append(lines, line)
	}

	file.AppendLineToFile(global.GetPathTableDataRecord(insert.Table), lines)
	if processUnique {
		file.AppendLineToFile(global.GetPathUniqueData(unique.Table, unique.Name), linesIndex)
	}

	general.RuntimeDone()
	c.JSON(200, "ok")

}
