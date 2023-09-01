package commands_func

import (
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
)

func getSequenceLastValue(sequenceName string) int {

	var sequenceJson command.Sequence

	err := file.LoadFileJSON(global.GetPathSequence(sequenceName), &sequenceJson)
	if err != nil {
		return -1
	}

	sequenceJson.LasValue = sequenceJson.LasValue + sequenceJson.IncrementValue
	SaveSequenceConfig(sequenceJson)

	return sequenceJson.LasValue
}

func CountLines(filename string) int {
	lines, _ := file.ReadLines(filename)
	return len(lines) - 1
}
