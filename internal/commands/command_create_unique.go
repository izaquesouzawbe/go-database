package commands

import (
	"fmt"
	"go-zdb-api/internal/global"
	"go-zdb-api/internal/models/command"
	"go-zdb-api/pkg/file"
)

func commandCreateUnique(query string) []map[string]string {

	fmt.Println(query)

	unique := command.Unique{
		Name:    extractUniqueName(query),
		Table:   extractUniqueTable(query),
		Columns: extractUniqueColumns(query),
	}

	path := global.GetPathUnique(unique.Table, unique.Name)
	if !file.FileExists(path) {
		file.CreateDir(path)
	}

	saveUniqueConfig(unique)
	saveUniqueData(unique)

	return nil
}
