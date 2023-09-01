package commands_func

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

func SaveSequenceConfig(sequence command.Sequence) {

	file := file.CreateFile(global.GetPathSequence(sequence.Name), true)
	saveFileJson(file, sequence)
}

func SaveDatabaseConfig(database command.Database) {

	file := file.CreateFile(global.GetPathConfigDatabase(), true)
	saveFileJson(file, database)

}

func SaveUniqueConfig(unique command.Unique) {

	file := file.CreateFile(global.GetPathUniqueConfig(unique.Table, unique.Name), true)
	saveFileJson(file, unique)

}

func SaveTableConfig(tableData command.Table) {

	file := file.CreateFile(global.GetPathTableConfig(tableData.TableName), true)
	saveFileJson(file, tableData)

}

func SaveUniqueData(unique command.Unique) {
	if !file.FileExists(global.GetPathUniqueData(unique.Table, unique.Name)) {
		file.CreateFile(global.GetPathUniqueData(unique.Table, unique.Name), false)
	}
}
