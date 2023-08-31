package commands

import (
	"fmt"
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
	saveSequenceConfig(sequenceJson)

	return sequenceJson.LasValue
}

func CountLines(filename string) int {
	lines, _ := file.ReadLines(filename)
	fmt.Println(filename)
	fmt.Println(len(lines))
	return len(lines)
}
