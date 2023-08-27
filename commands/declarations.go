package commands

type Table struct {
	TableName string              `json:"table_name"`
	Fields    []map[string]string `json:"fields"`
}

type Database struct {
	Name   string   `json:"name"`
	Tables []string `json:"tables"`
}

var LoadTablesVar []Table

var pathData string = "data/"
var pathDatabase string = "teste_db"
