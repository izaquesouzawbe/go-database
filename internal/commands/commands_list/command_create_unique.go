package commands_list

import (
	"fmt"
	"go-zdb-api/internal/commands/commands_func"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/dir"
	"go-zdb-api/pkg/file"
)

func CommandCreateUnique(query string) []map[string]string {

	fmt.Println(query)

	unique := command.Unique{
		Name:    commands_func.ExtractUniqueName(query),
		Table:   commands_func.ExtractUniqueTable(query),
		Columns: commands_func.ExtractUniqueColumns(query),
	}

	path := global.GetPathUnique(unique.Table, unique.Name)

	if !dir.DirExists(global.GetPathUniques(unique.Table)) {
		dir.CreateDir(global.GetPathUniques(unique.Table))
	}

	if !file.FileExists(path) {
		dir.CreateDir(path)
	}

	commands_func.SaveUniqueConfig(unique)
	commands_func.SaveUniqueData(unique)

	return nil
}
