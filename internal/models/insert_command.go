package models

type InsertCommand struct {
	TableName string
	Fields    []string
	Values    []string
}
