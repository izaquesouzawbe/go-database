package commands

var pathData string = "data/"
var pathDatabase string = "teste_db"

func setPathDatabase(value string) {
	pathDatabase = value
}

func getPathData() string {
	return pathData
}

func getPathDatabase() string {
	return getPathData() + pathDatabase + "/"
}

func getPathTable(table string) string {
	return getPathDatabase() + table + "/"
}

func getPathConfigTable(table string) string {
	return getPathTable(table) + "table.json"
}

func getPathFilesTable(table string) []string {

	newPath := getPathTable(table)

	var files []string
	files = append(files, newPath+"1.txt")
	return files
}
