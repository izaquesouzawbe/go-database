package commands

func commandInsertInto(query string) {

	/*insertCommand, _ := extractInsertCommandInfo(query)
	var tabelaJson Table

	err := file.LoadFileJSON(getPathConfigTable(insertCommand.TableName), &tabelaJson)
	if err != nil {
		return
	}

	line := ""

	fieldSequence := getSequenceOfFields(tabelaJson.Fields)
	var sequenceValue int

	/*if len(fieldSequence) > 0 {
		sequenceValue = getSequenceLastValue(tabelaJson.Sequence)
	}*/

	/*for _, field := range tabelaJson.Fields {

		if len(fieldSequence) > 0 {
			if field["name"] == fieldSequence {
				line += strconv.Itoa(sequenceValue) + ";"
			}
		}

		for iColumn, column := range insertCommand.Fields {

			if column == field["name"] {
				line += insertCommand.Values[iColumn] + ";"
			}

		}

	}

	file.AppendLineToFile(getPathTableDataRecord(insertCommand.TableName), line)*/

}
