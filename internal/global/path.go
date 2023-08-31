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
func GetPathIndexes(tableName string) string {
	return GetPathTable(tableName) + "indexes/"
}

func GetPathIndex(tableName, indexName string) string {
	return GetPathIndexes(tableName) + indexName + "/"
}

// UNIQUES
func GetPathUniques(tableName string) string {
	return GetPathTable(tableName) + "uniques/"
}

func GetPathUnique(tableName, indexName string) string {
	return GetPathUniques(tableName) + indexName + "/"
}

func GetPathUniqueConfig(tableName, indexName string) string {
	return GetPathUnique(tableName, indexName) + indexName + ".json"
}

func GetPathUniqueData(tableName, indexName string) string {
	return GetPathUnique(tableName, indexName) + "data.txt"
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
	return GetPathTableData(table) + "data.txt"
}
