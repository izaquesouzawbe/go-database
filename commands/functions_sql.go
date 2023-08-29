package commands

import "go-database/file"

func getSequenceLastValue(sequenceName string) int {

	var sequenceJson Sequence

	err := file.LoadFileJSON(getPathSequence(sequenceName), &sequenceJson)
	if err != nil {
		return -1
	}

	sequenceJson.LasValue = sequenceJson.LasValue + sequenceJson.IncrementValue
	saveSequenceConfig(sequenceJson)

	return sequenceJson.LasValue
}
