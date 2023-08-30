package commands

import (
	"go-zdb-api/internal/models"
	"go-zdb-api/pkg/file"
)

func getSequenceLastValue(sequenceName string) int {

	var sequenceJson models.Sequence

	err := file.LoadFileJSON(getPathSequence(sequenceName), &sequenceJson)
	if err != nil {
		return -1
	}

	sequenceJson.LasValue = sequenceJson.LasValue + sequenceJson.IncrementValue
	saveSequenceConfig(sequenceJson)

	return sequenceJson.LasValue
}
