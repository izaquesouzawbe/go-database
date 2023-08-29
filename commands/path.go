package commands

// DATA
func getPathData() string {
	return pathData
}

// DATABASE
func getPathDatabase() string {
	return getPathData() + pathDatabase + "/"
}

func setPathDatabase(value string) {
	pathDatabase = value
}

func getPathConfigDatabase() string {
	return getPathDatabase() + "database.json"
}

// SEQUENCE
func getPathSequences() string {
	return getPathDatabase() + "sequences/"
}

func getPathSequence(sequenceName string) string {
	return getPathSequences() + sequenceName + ".json"
}

// INDEXES
func getPathIndexes() string {
	return getPathDatabase() + "indexes/"
}

func getPathIndex(indexName string) string {
	return getPathIndexes() + indexName + ".bin"
}

// TABLES
func getPathTables() string {
	return getPathDatabase() + "tables/"
}

func getPathTable(table string) string {
	return getPathTables() + table + "/"
}

func getPathConfigTable(table string) string {
	return getPathTable(table) + "table.json"
}

func getPathTableData(table string) string {
	return getPathTable(table) + "data/"
}

func getPathTableDataRecord(table string) string {
	return getPathTableData(table) + "data1.bin"
}
