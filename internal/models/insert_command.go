package models

type InsertCommand struct {
	TableName string
	Fields    []string
}

type InsertCommandValues struct {
	Values []string
}
