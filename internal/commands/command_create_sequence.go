package commands

import (
	"go-zdb-api/internal/models/command"
)

func commandCreateSequence(commands []string) []string {

	msg := onValidateCreateSequence(commands)
	if len(msg) > 0 {
		return []string{msg}
	}

	sequenceName := commands[2]
	sequence := command.Sequence{
		Name:           sequenceName,
		LasValue:       1,
		IncrementValue: 1,
	}

	saveSequenceConfig(sequence)

	return []string{"Sequence '" + sequenceName + "' successfully created!"}

}

func onValidateCreateSequence(commands []string) string {
	return ""
}
