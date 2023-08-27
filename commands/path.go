package commands

func getPathTable(table string) string {
	return path + table + "/"
}

func getPathConfigTable(table string) string {
	return getPathTable(table) + "config.txt"
}

func getPathFilesTable(table string) []string {

	newPath := getPathTable(table)

	var files []string
	files = append(files, newPath+"1.txt")
	return files
}
