package commands

type Table struct {
	TableName string  `json:"table_name"`
	Fields    []Field `json:"fields"`
}

type Field struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	NotNull int    `json:"not_null"`
}

type Database struct {
	Name   string   `json:"name"`
	Tables []string `json:"tables"`
}

type Sequence struct {
	Name           string `json:"name"`
	LasValue       int    `json:"last_value"`
	IncrementValue int    `json:"increment_value"`
}

var LoadTablesVar []Table

var pathData string = "data/"
var pathDatabase string = "teste_db"
