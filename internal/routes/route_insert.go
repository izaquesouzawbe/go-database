package routes

import (
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/internal/models/request"
	"go-zdb-api/pkg/file"
	"go-zdb-api/pkg/general"
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

		lines = append(lines, line)
	}

	file.AppendLineToFile(global.GetPathTableDataRecord(insert.Table), lines)

	general.RuntimeDone()
	c.JSON(200, "ok")

}
