package commands

func getPathData() string {
	return pathData
}

func getPathDatabase() string {
	return getPathData() + pathDatabase + "/"
}

func setPathDatabase(value string) {
	pathDatabase = value
}

func getPathConfigDatabase() string {
	return getPathDatabase() + "database.json"
}

func getPathTable(table string) string {
	return getPathDatabase() + table + "/"
}

func getPathConfigTable(table string) string {
	return getPathTable(table) + "table.json"
}

func getPathDataTable(table string) string {
	return getPathTable(table) + "1.data"
}
