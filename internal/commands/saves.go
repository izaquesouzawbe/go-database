package commands

import (
	"encoding/json"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
	"io"
)

func saveFileJson(w io.Writer, v any) {
	encoder := json.NewEncoder(w)
	encoder.Encode(v)
}

func saveSequenceConfig(sequence command.Sequence) {

	file := file.CreateFile(global.GetPathSequence(sequence.Name), true)
	saveFileJson(file, sequence)
}

func saveDatabaseConfig(database command.Database) {

	file := file.CreateFile(global.GetPathConfigDatabase(), true)
	saveFileJson(file, database)

}

func saveUniqueConfig(unique command.Unique) {

	file := file.CreateFile(global.GetPathUniqueConfig(unique.Table, unique.Name), true)
	saveFileJson(file, unique)

}

func saveTableConfig(tableData command.Table) {

	file := file.CreateFile(global.GetPathTableConfig(tableData.TableName), true)
	saveFileJson(file, tableData)

}

func saveUniqueData(unique command.Unique) {
	file.CreateFile(global.GetPathUniqueData(unique.Table, unique.Name), false)
}
