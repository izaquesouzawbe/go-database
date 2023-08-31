package global

// DATA
func GetPathData() string {
	return pathData
}

// DATABASE
func GetPathDatabase() string {
	return GetPathData() + pathDatabase + "/"
}

func SetPathDatabase(value string) {
	pathDatabase = value
}

func GetPathConfigDatabase() string {
	return GetPathDatabase() + "database.json"
}

// SEQUENCE
func GetPathSequences() string {
	return GetPathDatabase() + "sequences/"
}

func GetPathSequence(sequenceName string) string {
	return GetPathSequences() + sequenceName + ".json"
}

// INDEXES
func GetPathIndexes() string {
	return GetPathDatabase() + "indexes/"
}

func GetPathIndex(indexName string) string {
	return GetPathIndexes() + indexName + ".bin"
}

// TABLES
func GetPathTables() string {
	return GetPathDatabase() + "tables/"
}

func GetPathTable(table string) string {
	return GetPathTables() + table + "/"
}

func GetPathTableConfig(table string) string {
	return GetPathTable(table) + "table.json"
}

func GetPathTableData(table string) string {
	return GetPathTable(table) + "data/"
}

func GetPathTableDataRecord(table string) string {
	return GetPathTableData(table) + "data1.bin"
}
